package main

import (
	"db"
	"encoding/json"
	"github.com/DeKugelschieber/go-session"
	"io/ioutil"
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

// Default configuration with server, port, database and log file.
type config struct {
	Host          string `json:"host"`
	TLSPrivateKey string `json:"tls_private_key"`
	TLSCert       string `json:"tls_cert"`
	DbUser        string `json:"dbuser"`
	DbPwd         string `json:"dbpwd"`
	DbHost        string `json:"dbhost"`
	Db            string `json:"db"`
	LogFile       string `json:"logfile"`
}

var (
	cfg config
)

// Loads config from json.
func loadConfig(file string) {
	cfg = config{}

	// read
	log.Print("Loading config file: ", file)
	content, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	// parse
	if err := json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}
}

// Log to file if logfile name is set.
func logToFile() *os.File {
	if cfg.LogFile == "" {
		return nil
	}

	if _, err := os.Stat(log_dir); err != nil {
		if err := os.Mkdir(log_dir, 0744); err != nil {
			panic(err)
		}
	}

	now := time.Now().Format("02.01.2006_15:04:05")
	handle, err := os.Create(filepath.Join(log_dir, cfg.LogFile+"_"+now))

	if err != nil {
		panic(err)
	}

	log.SetOutput(handle)

	return handle
}

// Starts the REST server.
func startServer() {
	log.Print("Starting server on ", cfg.Host)

	mux := http.NewServeMux()
	mux.Handle("/robots.txt", http.HandlerFunc(returnRobotsTxt))
	mux.Handle("/", http.FileServer(http.Dir(public_dir)))

	mux.HandleFunc("/api/checkLogin", http.HandlerFunc(rest.CheckSessionHandler))
	mux.HandleFunc("/api/login", http.HandlerFunc(rest.Login))
	mux.HandleFunc("/api/logout", http.HandlerFunc(rest.Logout))
	mux.Handle("/api/addEditUser", session.AccessMiddleware(http.HandlerFunc(rest.AddEditUser), returnSessionErr))
	mux.Handle("/api/removeUser", session.AccessMiddleware(http.HandlerFunc(rest.RemoveUser), returnSessionErr))
	mux.Handle("/api/getAllUsers", session.AccessMiddleware(http.HandlerFunc(rest.GetAllUser), returnSessionErr))
	mux.Handle("/api/getUser", session.AccessMiddleware(http.HandlerFunc(rest.GetUser), returnSessionErr))

	mux.Handle("/api/saveSettings", session.AccessMiddleware(http.HandlerFunc(rest.SaveSettings), returnSessionErr))
	mux.Handle("/api/getSettings", session.AccessMiddleware(http.HandlerFunc(rest.GetSettings), returnSessionErr))

	mux.Handle("/api/addEditConfiguration", session.AccessMiddleware(http.HandlerFunc(rest.AddEditConfiguration), returnSessionErr))
	mux.Handle("/api/removeConfiguration", session.AccessMiddleware(http.HandlerFunc(rest.RemoveConfiguration), returnSessionErr))
	mux.Handle("/api/getAllConfigurations", session.AccessMiddleware(http.HandlerFunc(rest.GetAllConfigurations), returnSessionErr))
	mux.Handle("/api/getConfiguration", session.AccessMiddleware(http.HandlerFunc(rest.GetConfiguration), returnSessionErr))

	mux.Handle("/api/startInstance", session.AccessMiddleware(http.HandlerFunc(rest.StartInstance), returnSessionErr))
	mux.Handle("/api/stopInstance", session.AccessMiddleware(http.HandlerFunc(rest.StopInstance), returnSessionErr))
	mux.Handle("/api/getAllInstances", session.AccessMiddleware(http.HandlerFunc(rest.GetAllInstances), returnSessionErr))
	mux.Handle("/api/getInstanceLog", session.AccessMiddleware(http.HandlerFunc(rest.GetInstanceLog), returnSessionErr))

	if cfg.TLSPrivateKey == "" || cfg.TLSCert == "" {
		if err := http.ListenAndServe(cfg.Host, mux); err != nil {
			panic(err)
		}
	} else {
		log.Print("Started with TLS enabled")

		if err := http.ListenAndServeTLS(cfg.Host, cfg.TLSCert, cfg.TLSPrivateKey, mux); err != nil {
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
	// load and setup log
	loadConfig(config_file)
	log := logToFile()

	if log != nil {
		defer log.Close()
	}

	// connect to db
	db.Connect(cfg.DbUser, cfg.DbPwd, cfg.DbHost, cfg.Db)
	defer db.Disconnect()

	// start session manager
	sessionProvider := session.NewMemProvider()
	session.New(session_cookie_name, session_cookie_lifetime, sessionProvider)

	// and go
	startServer()
}
