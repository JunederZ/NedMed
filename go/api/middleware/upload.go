package middleware

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type FileValidator struct {
	MaxSize    int64
	AllowedExt map[string]bool
}

func NewFileValidator() *FileValidator {
	return &FileValidator{
		MaxSize: 10 << 20, // 10MB default limit
		AllowedExt: map[string]bool{
			".jpg":  true,
			".jpeg": true,
			".png":  true,
			".gif":  true,
			".webp": true,
		},
	}
}

func (v *FileValidator) ValidateImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the file from the request
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No file uploaded",
			})
		}

		// Validate file size
		if file.Size > v.MaxSize {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "File too large",
			})
		}

		// Validate file extension
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !v.AllowedExt[ext] {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid file type",
			})
		}

		// Open the file to check its content
		fileHeader, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error processing file",
			})
		}
		defer func(fileHeader multipart.File) {
			err := fileHeader.Close()
			if err != nil {

			}
		}(fileHeader)

		// Read first 512 bytes to detect content type
		buff := make([]byte, 512)
		if _, err := fileHeader.Read(buff); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error reading file",
			})
		}

		// Validate content type
		filetype := http.DetectContentType(buff)
		if !strings.HasPrefix(filetype, "image/") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "File is not an image",
			})
		}

		// Store file in locals for the next handler
		c.Locals("file", file)

		return c.Next()
	}
}
