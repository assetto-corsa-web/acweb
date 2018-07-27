package session

import (
	"time"
)

// A provider abstracts the session backend used by the manager,
// so that different kinds of storages can be used.
// For an example see MemProvider.
type Provider interface {
	Init(string, time.Time) (Session, error)
	Read(string) (Session, error)
	Write(*Session) error
	Destroy(*Session) error
	GC()
}
