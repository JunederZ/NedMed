package routes

import (
	"NedMed/api/handler"
	"NedMed/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterFileRoutes(app *fiber.App) {
	validator := middleware.NewFileValidator()

	app.Post("/upload", validator.ValidateImage(), (&handler.FileHandler{}).UploadFile)
	app.Get("/files/:filename", (&handler.FileHandler{}).GetFile)
}
