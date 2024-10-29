package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	create = `
	CREATE TABLE IF NOT EXISTS article (
		id VARCHAR(255) NOT NULL,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		image VARCHAR(256) NOT NULL,
		PRIMARY KEY (id)
	);
	`
)

func NewDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASSWD"),
		Addr:   os.Getenv("DBADDR"),
		Net:    "tcp",
		DBName: "artic",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(create); err != nil {
		return nil, err
	}

	return db, nil
}