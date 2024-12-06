package service

import (
	_ "context"
	_ "errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"itfest/internal/models"
	"itfest/internal/repository"
	"itfest/internal/utils"
)

func CreateItem(conn *pgxpool.Conn, item models.Item, file []byte) (string, error) {
	imageKey := fmt.Sprintf("items/%s.jpg", item.Title)
	imageURL, err := utils.UploadFile(imageKey, file)
	if err != nil {
		return "", err
	}

	item.ImageURL = imageURL
	itemID, err := repository.CreateItem(conn, item)
	if err != nil {
		return "", err
	}

	return itemID, nil
}

func GetItemById(conn *pgxpool.Pool, id string) (*models.Item, error) {
	return repository.GetItemById(conn, id)
}
