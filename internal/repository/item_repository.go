package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"itfest/internal/models"
	"log"
)

func CreateItem(conn *pgxpool.Conn, item models.Item) (string, error) {
	query := `INSERT INTO items (title, description, image_url) VALUES ($1, $2, $3) RETURNING id`
	var itemID string
	err := conn.QueryRow(context.Background(), query, item.Title, item.Description, item.ImageURL).Scan(&itemID)
	if err != nil {
		log.Printf("Failed to create item: %v", err)
		return "", err
	}
	return itemID, nil
}

func GetItemById(conn *pgxpool.Pool, id string) (*models.Item, error) {
	var item models.Item
	query := `SELECT id, title, description, image_url FROM items WHERE id = $1`
	err := conn.QueryRow(context.Background(), query, id).Scan(&item.ID, &item.Title, &item.Description, &item.ImageURL)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
