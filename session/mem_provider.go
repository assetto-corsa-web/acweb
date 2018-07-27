package session

import (
	"errors"
	"sync"
	"time"
)

// MemProvider provides a in memory solution to store sessions.
// Use the NewMemProvider() function to create a new instance and pass
// it to Manager. The methods should not be called manually!
type MemProvider struct {
	Provider
	mutex    sync.Mutex
	sessions []Session
}

// Creates a new in memory provider for the session manager.
func NewMemProvider() *MemProvider {
	provider := MemProvider{}
	provider.sessions = make([]Session, 0)

	return &provider
}

// Stores a new session in memory with given token and lifetime.
func (p *MemProvider) Init(token string, lifetime time.Time) (Session, error) {
	session := emptySession()

	if token == "" {
		return session, errors.New("Token must be set")
	}

	if lifetime.Before(time.Now()) {
		return session, errors.New("Lifetime must be greater than now")
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, session := range p.sessions {
		if session.token == token {
			return session, errors.New("Session with token '" + token + "' exists already")
		}
	}

	session.token = token
	session.lifetime = lifetime
	p.sessions = append(p.sessions, session)

	return session, nil
}

// Reads a session from memory by given token.
func (p *MemProvider) Read(token string) (Session, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, session := range p.sessions {
		if session.token == token {
			return session, nil
		}
	}

	return emptySession(), errors.New("Session with token '" + token + "' not found")
}

// Updates a session in memory.
func (p *MemProvider) Write(session *Session) error {
	if session == nil {
		return errors.New("Session must not be nil")
	}

	if session.data == nil {
		return errors.New("Session must be initialized")
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i, s := range p.sessions {
		if s.token == session.token {
			p.sessions[i].data = session.data
			p.sessions[i].lifetime = session.lifetime
			return nil
		}
	}

	return errors.New("Session with token '" + session.token + "' not found")
}

// Removes a session from memory.
func (p *MemProvider) Destroy(session *Session) error {
	if session == nil {
		return errors.New("Session must not be nil")
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i, s := range p.sessions {
		if s.token == session.token {
			p.sessions = append(p.sessions[:i], p.sessions[i+1:]...)
			return nil
		}
	}

	return errors.New("Session with token '" + session.token + "' not found")
}

// Iterates over all existing sessions in memory and removes sessions which
// exceeded their lifetime.
func (p *MemProvider) GC() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	now := time.Now()
	removed := 0

	for i, session := range p.sessions {
		if session.lifetime.Before(now) {
			p.sessions = append(p.sessions[:i-removed], p.sessions[i-removed+1:]...)
			removed++
		}
	}
}
