package main

import (
	"NedMed/api/models/files"
	"NedMed/api/routes"
	"NedMed/internal/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := database.NewDatabase()

	if db == nil {
		log.Fatal("Error connecting to database")
	}
	//
	db.Migrate()

	// add seed if table empty
	if db.Conn.Migrator().HasTable(&files.FileEntity{}) {
		var count int64
		db.Conn.Model(&files.FileEntity{}).Count(&count)
		if count == 0 {
			db.Seed()
		}
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.RegisterFileRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
