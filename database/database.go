package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	OpenConns = 10
	IdleConns = 3
	LifeTime  = 60 * time.Second
)

func DbConnect() (*sql.DB, error) {
	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), PORT, os.Getenv("USER"), os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(OpenConns)
	DB.SetMaxIdleConns(IdleConns)
	DB.SetConnMaxLifetime(LifeTime)

	return DB, nil
}
