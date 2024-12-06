package service

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"itfest/internal/models"
	"itfest/internal/repository"
	"itfest/internal/utils"
)

func CreateItem(conn *pgxpool.Conn, item models.Item, file []byte) (int, error) {
	imageKey := fmt.Sprintf("items/%s.jpg", item.Title)
	imageURL, err := utils.UploadFile(imageKey, file)
	if err != nil {
		return 0, err
	}

	item.ImageURL = imageURL
	itemID, err := repository.CreateItem(conn, item)
	if err != nil {
		return 0, err
	}

	return itemID, nil
}

func GetItemById(conn *pgxpool.Pool, id int) (*models.Item, error) {
	return repository.GetItemById(conn, id)
}
