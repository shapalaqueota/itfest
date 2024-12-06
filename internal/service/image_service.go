package service

import (
	"fmt"
	"github.com/google/uuid"
	"itfest/internal/utils"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type ImageService struct {
	allowedTypes map[string]bool
	maxFileSize  int64
}

func NewImageService() *ImageService {
	return &ImageService{
		allowedTypes: map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/gif":  true,
		},
		maxFileSize: 5 << 20, // 5MB
	}
}

func (s *ImageService) UploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Validate file size
	if header.Size > s.maxFileSize {
		return "", fmt.Errorf("file size exceeds maximum limit")
	}

	// Read file content
	buffer := make([]byte, header.Size)
	if _, err := file.Read(buffer); err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	// Detect content type
	contentType := http.DetectContentType(buffer)
	if !s.allowedTypes[contentType] {
		return "", fmt.Errorf("invalid file type: %s", contentType)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Upload to S3
	key, err := utils.UploadFile(filename, buffer)
	if err != nil {
		return "", err
	}

	return key, nil
}
