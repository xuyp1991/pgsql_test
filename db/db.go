package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/cihub/seelog"
)

var db *sql.DB

func Init() {
	var err error

	if db == nil {
		pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"127.0.0.1","5432","xuyp","123456","testdb")
			//Config.DB.IP, Config.DB.Port, Config.DB.User, Config.DB.Passwd, Config.DB.Database)

		db, err = sql.Open("postgres", pgInfo)

		if err != nil {
			seelog.Errorf("open db failed, err: %v", err)
		}
	}

	err = db.Ping()

	if err != nil {
		seelog.Errorf("connect to db failed, err: %v", err)
	}

	seelog.Tracef("connect to db succ, db: %v", db)
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		seelog.Errorf("close db err, db: %v, err: %v", db, err)
	}
	seelog.Tracef("close db succ")
}
