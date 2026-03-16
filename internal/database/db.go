package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./novel.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS novels (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			slug TEXT UNIQUE,
			source_url TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS chapters (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			novel_id INTEGER,
			title TEXT,
			slug TEXT,
			source_url TEXT UNIQUE,
			content_raw TEXT,
			content_translated TEXT,
			status TEXT DEFAULT 'pending',
			FOREIGN KEY(novel_id) REFERENCES novels(id)
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
