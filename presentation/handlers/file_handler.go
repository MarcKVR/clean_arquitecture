package handlers

import (
	"github.com/MarcKVR/clean_arquitecture/application/service"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	service *service.FileService
}

func NewFileHandler(service *service.FileService) *FileHandler {
	return &FileHandler{
		service: service,
	}
}

func (h *FileHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file found",
		})
	}

	url, err := h.service.UploadFile(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": url,
	})
}
