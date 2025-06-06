package util

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const UploadDir = "uploads"

// const DefaultImage = "uploads/default_image.jpeg"
var DefaultImage = filepath.Join(UploadDir, "default_image.jpeg")

func SaveImage(c *fiber.Ctx) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return DefaultImage, err
	}

	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		os.Mkdir(UploadDir, os.ModePerm)
	}

	filePath := filepath.Join(UploadDir, file.Filename)
	err = c.SaveFile(file, filePath)
	if err != nil {
		return DefaultImage, err
	}
	return filePath, nil
}

func DeleteImage(filePath string) error {
	if filePath == DefaultImage {
		return nil
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(filePath)
}
