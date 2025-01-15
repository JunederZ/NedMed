package handler

import (
	"NedMed/api/models/files"
	"fmt"
	"mime/multipart"
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

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	fileExt := strings.Split(file.Filename, ".")[1]

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not save file",
		})
	}

	imageUrl := fmt.Sprintf("http://localhost:3000/files/%s", image)

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
	return c.SendFile(fmt.Sprintf("./uploads/%s", filename))
}
