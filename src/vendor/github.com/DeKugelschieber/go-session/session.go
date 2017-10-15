package session

import (
	"errors"
	"net/http"
	"reflect"
	"time"
)

// Session provides access to user context related data.
// Generic data can be set and obtained and will be available
// on page switch. Call Save() to make sure the data is stored.
// A session must be initialized by the session manager, to make
// sure it can be used later.
type Session struct {
	token    string
	lifetime time.Time
	data     map[string]interface{}
}

// Saves the session. Call when data was changed using the Set() method
// to make it available on next call.
func (s *Session) Save() error {
	if m == nil {
		return errors.New("Sessions not started")
	}

	return m.updateSession(s)
}

// Destroys and invalidates the session.
// http.ResponseWriter and http.Request are optional, if both are set,
// the session cookie will be destroyed (which is recommended).
func (s *Session) Destroy(w http.ResponseWriter, r *http.Request) error {
	if m == nil {
		return errors.New("Sessions not started")
	}

	err := m.destroySession(w, r, s)
	s.token = ""

	return err
}

// Resets the lifetime of this session.
func (s *Session) Renew() error {
	s.lifetime = time.Now().Add(m.maxLifetime)
	return m.updateSession(s)
}

// Returns the session token, which is a unique base 64 URL encoded string.
func (s *Session) Token() string {
	return s.token
}

// Returns true when this session is active.
func (s *Session) Active() bool {
	return s.token != ""
}

// Stores session variable for given key within given variable.
// When the variable cannot be found, is not assignable (not a pointer)
// or is nil, an error will be returned.
// If the type of of the variable found does not match the type of the
// variable it will be stored in, this method will panic.
func (s *Session) Get(key string, value interface{}) error {
	if value == nil {
		return errors.New("Value must not be nil for key '" + key + "'")
	}

	v, ok := s.data[key]

	if !ok {
		return errors.New("Value for key '" + key + "' not found")
	}

	reflectValue := reflect.ValueOf(value)
	kind := reflectValue.Kind()

	if kind != reflect.Ptr && kind != reflect.Interface {
		return errors.New("Value type invalid for key '" + key + "' (see https://golang.org/pkg/reflect/#Value.Elem)")
	}

	elem := reflectValue.Elem()

	if !elem.CanSet() {
		return errors.New("Cannot assign value for key '" + key + "' (see https://golang.org/pkg/reflect/#Value.Set)")
	}

	// panics
	elem.Set(reflect.ValueOf(v))

	return nil
}

// Sets session variable for given key. If a variable is stored for that key already,
// it will be replaced.
// Use the Save() method to store it for later usage and obtain it with Get().
func (s *Session) Set(key string, value interface{}) error {
	if value == nil {
		return errors.New("Value must not be nil for key '" + key + "'")
	}

	s.data[key] = value
	return nil
}

// Removes session variable for given key.
// Use the Save() method to remove it persistently.
func (s *Session) Remove(key string) {
	delete(s.data, key)
}
