package service

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"itfest/internal/models"
	"itfest/internal/repository"
	"itfest/internal/utils"
)

func CreateUser(conn *pgxpool.Conn, user models.User) (string, error) {
	err := userDataValidation(conn, user)
	userID, err := repository.CreateUser(conn, user)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func Login(conn *pgxpool.Pool, email, password string) (*models.User, error) {
	user, err := repository.GetUserByEmail(conn, email)
	if err != nil {
		return nil, err
	}

	passwordIsValid := utils.CheckPasswordHash(password, user.Password)

	if !passwordIsValid {
		return nil, errors.New("credentials invalid")
	}

	return user, nil
}

func GetUserById(conn *pgxpool.Pool, uuid string) (*models.User, error) {
	return repository.GetUserById(conn, uuid)
}

func userDataValidation(conn *pgxpool.Conn, user models.User) error {
	exists, err := repository.CheckUserExists(conn, user.Email, user.PhoneNumber)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user with this email or phone number already exists")
	}
	return nil
}
