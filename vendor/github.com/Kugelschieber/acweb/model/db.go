package model

import (
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	default_db_type = "mysql"
	connect_timeout = "60"
)

var (
	session *sqlx.DB
)

// Connects to database. Will panic on error.
// Pass host and database to use.
func Connect(user, password, host, database string) {
	// check db type is supported
	dbtype := GetDBType()

	if dbtype != "mysql" && dbtype != "postgres" {
		log.WithFields(log.Fields{"db_type": dbtype}).Fatal("Unknown database type provided, allowed are: mysql, postgres")
	}

	// connect
	log.Info("Connecting to '" + GetDBType() + "' database '" + database + "' at '" + host + "'")
	var s *sqlx.DB
	var err error

	if dbtype == "mysql" {
		s, err = sqlx.Connect("mysql", mysqlConnection(user, password, host, database))
	} else {
		s, err = sqlx.Connect("postgres", postgresConnection(user, password, host, database))
	}

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error connecting to database")
	}

	if err := s.Ping(); err != nil {
		log.WithFields(log.Fields{"err": err}).Fatal("Error pinging database")
	}

	session = s
}

func mysqlConnection(user, password, host, database string) string {
	return user + ":" + password + "@" + host + "/" + database
}

func postgresConnection(user, password, host, database string) string {
	return "host=" + host +
		" port=" + os.Getenv("ACWEB_DB_PORT") +
		" user=" + user +
		" password=" + password +
		" dbname=" + database +
		" sslmode=" + os.Getenv("ACWEB_DB_SSLMODE") +
		" sslcert=" + os.Getenv("ACWEB_DB_SSL_CERT") +
		" sslkey=" + os.Getenv("ACWEB_DB_SSL_KEY") +
		" sslrootcert=" + os.Getenv("ACWEB_DB_ROOT_CERT") +
		" connect_timeout=" + connect_timeout
}

// Disconnect from database. Call on shutdown!
func Disconnect() {
	session.Close()
}

// Returns the database handle.
func Get() *sqlx.DB {
	return session
}

// Returns the database type used in lower case.
func GetDBType() string {
	dbtype := os.Getenv("ACWEB_DB_TYPE")

	if dbtype == "" {
		return default_db_type
	}

	return strings.ToLower(os.Getenv("ACWEB_DB_TYPE"))
}
