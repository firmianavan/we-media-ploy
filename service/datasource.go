package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var db *sql.DB
var once sync.Once
var connStr string = "van:123456@tcp(127.0.0.1:3306)/wmp?parseTime=true&loc=UTC"

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatal("program exit: failed to connect the database %s", connStr)
		}
	})
	return db
}
