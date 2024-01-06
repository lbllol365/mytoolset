package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	// TODO DB自定义配置
	db, err := sql.Open("sqlite3", "foo.db")
	if err != nil {
		log.Fatal("DB初始化失败")
	}
	DB = db
}
