package storage

import "web3/internal/models"

type Storage interface {
	NewUserAuth(userID int, login, passwordHash string) error
	InsertUserData(userData models.UserData) (int, error)
	UpdateUserData(userID string, userData models.UserData) error
	GetUserByID(userID string) (models.UserData, error)
	GetPasswordHash(login string) (string, string, error)
	Close()
}
