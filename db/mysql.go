package db

import (
	"PLAYLISTBOOK/config"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {
	config := config.Get()
	dsn := fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true",
		config.DbUsername, config.DbPassword, config.DbHost,
		config.DbPort, config.DbName,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("connection error")
	}

	db.SetConnMaxLifetime(time.Minute)

	return db
}
