package db

import (
	"changeme/backend/config"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", config.Config.DatabaseConfig.Location)
	if err != nil {
		log.Fatal("DB初始化失败")
	}
	DB = db
}
