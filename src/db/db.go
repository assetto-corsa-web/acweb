package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	session *sql.DB
)

// Connects to database. Will panic on error.
// Pass host and database to use.
func Connect(user, password, host, database string) {
	// connect
	log.Print("Connecting to database '" + database + "' at '" + host + "'")
	s, _ := sql.Open("mysql", user+":"+password+"@"+host+"/"+database)

	if err := s.Ping(); err != nil {
		log.Fatal(err)
	}

	session = s
}

// Disconnect from database. Call on shutdown!
func Disconnect() {
	session.Close()
}

// Returns the database handle.
func Get() *sql.DB {
	return session
}
