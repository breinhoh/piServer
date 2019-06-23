package database

import (
	"fmt"
	"github.com/rhinoman/couchdb-go"
	"os"
	"piServer/configs"
	"time"
)

// Db for use in store file
var Db *couchdb.Database

// Init starts the db connection
func Init() {

	dbConfig := config.GetDBConfig()
	var timeout = time.Duration(time.Duration(dbConfig.Timeout) * time.Millisecond)
	conn, _ := couchdb.NewConnection(dbConfig.IP, dbConfig.Port, timeout)
	auth := couchdb.BasicAuth{Username: dbConfig.Username, Password: dbConfig.Password}
	Db = conn.SelectDB(dbConfig.Name, &auth)
	if err := Db.DbExists; err != nil {
		fmt.Printf("Could not connect to database.\n")
		os.Exit(1)
	}

}

// GetDb returns the db for use
func GetDb() *couchdb.Database {
	return Db
}
