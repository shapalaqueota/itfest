package api

import (
	"itfest/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadImageHandler handles the image upload
func UploadImageHandler(c *gin.Context) {
	// Retrieve the file from the request
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file: " + err.Error()})
		return
	}
	defer file.Close()

	// Create an instance of ImageService
	imageService := service.NewImageService()

	// Upload the image using ImageService
	fileKey, err := imageService.UploadImage(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with the file key or URL
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "fileKey": fileKey})
}
