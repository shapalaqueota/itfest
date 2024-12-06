package api

import (
	"context"
	"itfest/internal/db"
	"itfest/internal/models"
	"itfest/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItemHandler(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if item.Title == "" || item.Description == "" || item.Price == "" || item.Category == "" || item.ImageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required (title, description, price, category, image_url)"})
		return
	}

	conn, err := db.DB.Acquire(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire database connection"})
		return
	}
	defer conn.Release()

	itemID, err := repository.CreateItem(conn, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item: " + err.Error()})
		return
	}

	// Respond with the created item's ID
	c.JSON(http.StatusOK, gin.H{
		"message": "Item created successfully",
		"itemID":  itemID,
	})
}
