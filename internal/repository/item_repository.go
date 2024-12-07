package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"itfest/internal/models"
	"log"
)

func CreateItem(conn *pgxpool.Conn, item models.Item) (int, error) {
	query := `INSERT INTO item (title, description, price, category, image_url) 
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id`
	var itemID int
	err := conn.QueryRow(context.Background(), query,
		item.Title, item.Description, item.Price, item.Category, item.ImageURL).Scan(&itemID)
	if err != nil {
		log.Printf("Failed to create item: %v", err)
		return 0, err
	}
	return itemID, nil
}

func GetItemById(conn *pgxpool.Pool, id int) (*models.Item, error) {
	query := `SELECT id, title, description, price, category, image_url FROM item WHERE id = $1`
	var item models.Item
	err := conn.QueryRow(context.Background(), query, id).Scan(
		&item.ID,
		&item.Title,
		&item.Description,
		&item.Price,
		&item.Category,
		&item.ImageURL,
	)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func DeleteItem(conn *pgxpool.Conn, id int) error {
	query := `DELETE FROM item WHERE id = $1 RETURNING id`
	var deletedID int
	err := conn.QueryRow(context.Background(), query, id).Scan(&deletedID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllItems(conn *pgxpool.Conn) ([]models.Item, error) {
	query := `SELECT id, title, description, price, category, image_url FROM item`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Category, &item.ImageURL)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
