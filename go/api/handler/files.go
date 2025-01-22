package handler

import (
	"NedMed/api/models/files"
	"NedMed/internal/database"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FileHandler struct {
	uploadDir string
}

func NewFileHandler(uploadDir string) *FileHandler {
	return &FileHandler{uploadDir: uploadDir}
}

func (h *FileHandler) UploadFile(c *fiber.Ctx) error {

	var req files.UploadRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(files.ErrorResponse{
			Success: false,
			Error:   "Invalid request format",
		})
	}

	file, ok := c.Locals("file").(*multipart.FileHeader)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving file",
		})
	}

	title, ok := c.Locals("title").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid title",
		})
	}

	description, ok := c.Locals("description").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid description",
		})
	}

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save to database

	err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error uploading file",
		})
	}

	imageUrl := fmt.Sprintf("http://localhost:3000/files/%s", image)

	fileEntity := files.FileEntity{
		Filename:    title,
		Size:        file.Size,
		MimeType:    file.Header["Content-Type"][0],
		Description: description,
		UploadedAt:  time.Now(),
		Url:         imageUrl,
	}

	db := database.NewDatabase()
	err = db.Conn.Create(&fileEntity).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not save file",
		})
	}

	fileInfo := files.FileResponse{
		ID:          uuid.New().String(),
		Filename:    file.Filename,
		Size:        file.Size,
		MimeType:    file.Header.Get("Content-Type"),
		Description: req.Description,
		UploadedAt:  time.Now(),
		URL:         imageUrl,
	}

	return c.JSON(files.UploadResponse{
		Success:  true,
		Message:  "File uploaded successfully",
		FileInfo: fileInfo,
	})
}

func (h *FileHandler) GetFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	fmt.Println(filename)
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		c.Set("Content-Type", "image/jpeg")
	case ".png":
		c.Set("Content-Type", "image/png")
	}
	return c.SendFile(fmt.Sprintf("./uploads/%s", filename))
}

func (h *FileHandler) GetAllFile(c *fiber.Ctx) error {
	print("Get all files")
	db := database.NewDatabase()
	print("Get all files [2]")
	var files []files.FileEntity
	err := db.Conn.Find(&files).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not retrieve files",
		})
	}

	return c.JSON(files)
}
