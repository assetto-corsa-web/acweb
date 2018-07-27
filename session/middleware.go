package session

import (
	"net/http"
)

// Redirect function in case the session is not set.
// If true is returned, the middleware will continue calling the next http.Handler.
type RedirectFunc func(http.ResponseWriter, *http.Request) bool

// Handler function with session injected.
type HandlerFunc func(Session, http.ResponseWriter, *http.Request)

// Middleware to check if session is set.
// Pass the next http.Handler to be called and a redirect function,
// which is called when session is not set.
// If you pass nil for the redirect function, it will return nothing to the client.
func AccessMiddleware(next http.Handler, redirect RedirectFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetCurrentSession(r)

		if !session.Active() && redirect != nil && !redirect(w, r) {
			return
		}

		// go on
		next.ServeHTTP(w, r)
	})
}

// Middleware to inject current session.
func Middleware(next HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := GetCurrentSession(r)
		next(session, w, r)
	})
}
