package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() *Database {
	db := &Database{}
	db.Connect()
	return db
}

func (db *Database) Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	db.Conn = conn
}

func (db *Database) Close() {
	err := db.Conn.Close()
	if err != nil {
		return
	}
}

func (db *Database) Migrate() {
	_, err := db.Conn.Exec(`
		CREATE TABLE IF NOT EXISTS files (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255),
			size INT,
			mime_type VARCHAR(255),
			description TEXT,
			uploaded_at TIMESTAMP,
			url VARCHAR(255)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Seed() {
	_, err := db.Conn.Exec(`
		INSERT INTO files (filename, size, mime_type, description, uploaded_at, url)
		VALUES (
			'example.jpg',
			1024,
			'image/jpeg',
			'Example image',
			NOW(),
			'http://localhost:3000/files/example.jpg'
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Drop() {
	_, err := db.Conn.Exec("DROP TABLE IF EXISTS files")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Reset() {
	db.Drop()
	db.Migrate()
	db.Seed()
}
