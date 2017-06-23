# go-session

Session management for Go.

## Features

* easy to learn and use
* tiny API
* custom session providers to store sessions in any database or memory structure you want
* comes with an in-memory provider
* concurrency safe
* store and receive any type within a session easily
* middleware to protect pages
* middleware to inject current session

## Installation

```
go get github.com/DeKugelschieber/go-session
```

## Usage

Here is a basic example on how to use this library. For full documentation please see below.

```
import s "github.com/DeKugelschieber/go-session"

// must be called on startup, cookie name, max lifetime and provider
err := s.New("cookie_name", 3600, NewMemProvider())

// create a new session using http.ResponseWriter and http.Request
session, err := s.NewSession(w, r)

// store data within session
err := session.Set("myobj", &myObj)
err := session.Save()

// get data from session
var myObj MyObj
err := session.Get("myobj", &myObj)

// remove data from session
session.Remove("myobj")

// destroy the session
err := session.Destroy(w, r)
```

## Middleware

There are two types of middleware. The first one protects endpoints by checking if the session is set. The second one injects the session into the handler, even if not set (use Session.Active() to check if the session is initialized):

```
import s "github.com/DeKugelschieber/go-session"

func redirect() bool {
    return true // true -> call next handler
}

mux.Handle("/route", s.AccessMiddleware(http.HandlerFunc(nextHandler), redirect))
```

```
func nextHandler(session s.Session, w http.ResponseWriter, r *http.Request) {
    // ...
}

mux.Handle("/route", s.Middleware(nextHandler))
```

## Custom providers

To use a custom provider, you have to implement the Provider interface and pass it when calling *New()*. The Provider interface is defined as:

```
type Provider interface {
    Init(string, time.Time) (Session, error)
    Read(string) (Session, error)
    Write(*Session) error
    Destroy(*Session) error
    GC()
}
```

For an example, see the MemProvider.

## Documentation

For full documentation please visit https://godoc.org/github.com/DeKugelschieber/go-session.

## Contribute

Contribution is welcome. Please create a pull request describing what you improved or create an issue to report bugs or discuss ideas.

## License

MIT
