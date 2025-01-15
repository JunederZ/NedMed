package database

import (
	"NedMed/api/models/files"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase() *Database {
	db := &Database{}
	db.Connect()
	return db
}

func (db *Database) Connect() {
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Conn = conn

}

func (db *Database) Close() {
	if db.Conn != nil {
		db.Conn = nil
	}
}

func (db *Database) Migrate() {
	err := db.Conn.AutoMigrate(&files.FileEntity{})
	if err != nil {
		log.Fatal(err)
	}

}

func (db *Database) Seed() {
	file := files.FileEntity{
		Filename:    "example.jpg",
		Size:        1024,
		MimeType:    "image/jpeg",
		Description: "Example image",
		UploadedAt:  time.Now(),
		Url:         "http://localhost:3000/files/example.jpg",
	}

	err := db.Conn.Create(&file).Error
	if err != nil {
		log.Fatal(err)
	}

}

func (db *Database) Drop() {
	err := db.Conn.Migrator().DropTable(&files.FileEntity{})
	if err != nil {
		log.Fatal(err)
	}

}

func (db *Database) Reset() {
	db.Drop()
	db.Migrate()
	db.Seed()
}
