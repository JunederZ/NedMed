package main

import (
	"NedMed/api/routes"
	"NedMed/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := database.NewDatabase()

	db.Seed()

	app := fiber.New()

	routes.RegisterFileRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
