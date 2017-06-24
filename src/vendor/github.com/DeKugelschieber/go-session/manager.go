// Package session provides simple server side session handling.
package session

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strconv"
	"time"
)

var (
	m *manager
)

type manager struct {
	cookieName  string
	maxLifetime time.Duration
	provider    Provider
}

// Creates a new session manager with given cookie name and max lifetime in seconds.
// The cookie name must not be an emtpy string, the lifetime must be greater than 0
// and a provider must be passed. If not, an error will be returned.
func New(cookieName string, maxLifetime int, provider Provider) error {
	if cookieName == "" {
		return errors.New("The cookie name must be set")
	}

	if maxLifetime <= 0 {
		return errors.New("The maximum lifetime of cookies must be greater than 0")
	}

	if provider == nil {
		return errors.New("The provider must be set")
	}

	m = &manager{cookieName, time.Duration(maxLifetime) * time.Second, provider}
	m.gc()

	return nil
}

// Creates a new session and returns it.
// If the http.ResponseWriter or http.Request is not set, the session token
// won't be stored in cookie. So to create a new session without using http,
// call it with both nil as parameters.
func NewSession(w http.ResponseWriter, r *http.Request) (Session, error) {
	if m == nil {
		return emptySession(), errors.New("Sessions not started")
	}

	token := generateSessionToken(m.cookieName)
	lifetime := time.Now().Add(m.maxLifetime)
	session, err := m.provider.Init(token, lifetime)

	if err != nil {
		return emptySession(), err
	}

	session.token = token
	session.lifetime = lifetime
	session.data = make(map[string]interface{})

	// init cookie if request is set
	if w != nil && r != nil {
		newCookie := http.Cookie{Name: m.cookieName,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(m.maxLifetime.Seconds())}
		http.SetCookie(w, &newCookie)
	}

	return session, nil
}

// Returns a session by session token if found.
// If not, an empty session and an error will be returned.
func GetSession(token string) (Session, error) {
	if m == nil {
		return emptySession(), errors.New("Sessions not started")
	}

	session, err := m.provider.Read(token)

	if err != nil {
		return emptySession(), err
	}

	session.lifetime = time.Now().Add(m.maxLifetime) // reset lifetime
	err = m.updateSession(&session)

	if err != nil {
		return emptySession(), err
	}

	return session, nil
}

// Returns the session for the http.Request.
// If not found, an empty session and an error will be returned.
func GetCurrentSession(r *http.Request) (Session, error) {
	if m == nil {
		return emptySession(), errors.New("Sessions not started")
	}

	if r == nil {
		return emptySession(), errors.New("Request must be set to obtain the current session")
	}

	cookie, err := r.Cookie(m.cookieName)

	if err != nil || cookie.Value == "" {
		return emptySession(), errors.New("Session cookie not set")
	}

	return GetSession(cookie.Value)
}

func (m *manager) updateSession(session *Session) error {
	return m.provider.Write(session)
}

func (m *manager) destroySession(w http.ResponseWriter, r *http.Request, session *Session) error {
	// unset cookie
	if w != nil && r != nil {
		cookie := http.Cookie{Name: m.cookieName, Path: "/", HttpOnly: true, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}

	return m.provider.Destroy(session)
}

func (m *manager) gc() {
	m.provider.GC()
	time.AfterFunc(m.maxLifetime, func() { m.gc() })
}

func generateSessionToken(seed string) string {
	now := strconv.Itoa(time.Now().Nanosecond())
	hash := sha256.New()
	hash.Write([]byte(seed + now))
	token := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return token
}

func emptySession() Session {
	session := Session{}
	session.data = make(map[string]interface{})
	return session
}
