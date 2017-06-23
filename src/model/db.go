package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var (
	session *sqlx.DB
)

// Connects to database. Will panic on error.
// Pass host and database to use.
func Connect(user, password, host, database string) {
	// connect
	log.Info("Connecting to database '" + database + "' at '" + host + "'")
	s, err := sqlx.Connect("mysql", user+":"+password+"@"+host+"/"+database)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error connecting to database")
	}

	if err := s.Ping(); err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error pinging database")
	}

	session = s
}

// Disconnect from database. Call on shutdown!
func Disconnect() {
	session.Close()
}

// Returns the database handle.
func Get() *sqlx.DB {
	return session
}
