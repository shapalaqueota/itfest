package api

import (
	"context"
	"itfest/internal/db"
	"itfest/internal/models"
	"itfest/internal/repository"
	"itfest/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateItemHandler(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	price := c.PostForm("price")
	category := c.PostForm("category")

	if title == "" || description == "" || price == "" || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required (title, description, price, category)"})
		return
	}

	header, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file: " + err.Error()})
		return
	}

	file, err := header.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image file: " + err.Error()})
		return
	}
	defer file.Close()

	imageService := service.NewImageService()

	imageURL, err := imageService.UploadImage(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed: " + err.Error()})
		return
	}

	conn, err := db.DB.Acquire(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire database connection"})
		return
	}
	defer conn.Release()

	item := models.Item{
		Title:       title,
		Description: description,
		Price:       price,
		Category:    category,
		ImageURL:    imageURL,
	}

	itemID, err := repository.CreateItem(conn, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Item created successfully",
		"itemID":   itemID,
		"imageURL": imageURL,
	})
}

func GetItemByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item id"})
		return
	}

	item, err := repository.GetItemById(db.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          item.ID,
		"title":       item.Title,
		"description": item.Description,
		"price":       item.Price,
		"category":    item.Category,
		"image_url":   item.ImageURL,
	})
}
