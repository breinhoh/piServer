package database

import (
	"github.com/rhinoman/couchdb-go"
	"piServer/configs"
	"time"
)

var Db *couchdb.Database

func Init() {

	dbConfig := config.GetDBConfig()
	var timeout = time.Duration(time.Duration(dbConfig.Timeout) * time.Millisecond)
	conn, _ := couchdb.NewConnection(dbConfig.IP, dbConfig.Port, timeout)
	auth := couchdb.BasicAuth{Username: dbConfig.Username, Password: dbConfig.Password}
	Db = conn.SelectDB(dbConfig.Name, &auth)
}

func GetDb() *couchdb.Database {
	return Db
}
