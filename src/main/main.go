package main

import (
	"db"
	"github.com/DeKugelschieber/go-session"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"rest"
	"time"
)

const (
	config_file = "config.json"
	log_dir     = "log"
	public_dir  = "public"
	robots_txt  = "robots.txt"

	session_cookie_name     = "acweb-session"
	session_cookie_lifetime = 60 * 60 * 24
)

// Log to file if logfile name is set.
func logToFile() *os.File {
	if os.Getenv("ACWEB_LOGDIR") == "" {
		return nil
	}

	if _, err := os.Stat(log_dir); err != nil {
		if err := os.Mkdir(log_dir, 0744); err != nil {
			panic(err)
		}
	}

	now := time.Now().Format("02.01.2006_15:04:05")
	handle, err := os.Create(filepath.Join(log_dir, os.Getenv("ACWEB_LOGDIR")+"_"+now))

	if err != nil {
		panic(err)
	}

	log.SetOutput(handle)

	return handle
}

// Starts the RESTful server.
func startServer() {
	log.Print("Starting server on ", os.Getenv("ACWEB_HOST"))

	mux := http.NewServeMux()
	mux.Handle("/robots.txt", http.HandlerFunc(returnRobotsTxt))
	mux.Handle("/", http.FileServer(http.Dir(public_dir)))
	mux.HandleFunc("/api/session", http.HandlerFunc(rest.CheckSession))
	mux.HandleFunc("/api/login", http.HandlerFunc(rest.Login))
	mux.HandleFunc("/api/logout", http.HandlerFunc(rest.Logout))
	mux.Handle("/api/user", session.AccessMiddleware(http.HandlerFunc(rest.UserHandler), returnSessionErr))
	mux.Handle("/api/settings", session.AccessMiddleware(http.HandlerFunc(rest.SettingsHandler), returnSessionErr))
	mux.Handle("/api/configuration", session.AccessMiddleware(http.HandlerFunc(rest.ConfigurationHandler), returnSessionErr))
	mux.Handle("/api/tracks", session.AccessMiddleware(http.HandlerFunc(rest.GetAvailableTracks), returnSessionErr))
	mux.Handle("/api/cars", session.AccessMiddleware(http.HandlerFunc(rest.GetAvailableCars), returnSessionErr))
	mux.Handle("/api/instance", session.AccessMiddleware(http.HandlerFunc(rest.InstanceHandler), returnSessionErr))
	mux.Handle("/api/instance/log", session.AccessMiddleware(http.HandlerFunc(rest.InstanceLogHandler), returnSessionErr))

	if os.Getenv("ACWEB_TLS_PRIVATE_KEY") == "" || os.Getenv("ACWEB_TLS_CERT") == "" {
		if err := http.ListenAndServe(os.Getenv("ACWEB_HOST"), mux); err != nil {
			panic(err)
		}
	} else {
		log.Print("Started with TLS enabled")

		if err := http.ListenAndServeTLS(os.Getenv("ACWEB_HOST"), os.Getenv("ACWEB_TLS_CERT"), os.Getenv("ACWEB_TLS_PRIVATE_KEY"), mux); err != nil {
			panic(err)
		}
	}
}

// Returns the robots.txt if exists.
func returnRobotsTxt(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, robots_txt)
}

// Returns an error to client on access to protected area when user session is not set.
func returnSessionErr(w http.ResponseWriter, r *http.Request) bool {
	s, _ := session.GetCurrentSession(r)
	return s.Active()
}

func main() {
	// setup log
	log := logToFile()

	if log != nil {
		defer log.Close()
	}

	// connect to db
	db.Connect(os.Getenv("ACWEB_DB_USER"), os.Getenv("ACWEB_DB_PASSWORD"), os.Getenv("ACWEB_DB_HOST"), os.Getenv("ACWEB_DB"))
	defer db.Disconnect()

	// start session manager
	sessionProvider := session.NewMemProvider()
	session.New(session_cookie_name, session_cookie_lifetime, sessionProvider)

	// and go
	startServer()
}
