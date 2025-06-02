package storage

import "web3/internal/models"

type Storage interface {
	NewUserAuth(userID int, login, passwordHash string) error
	InsertUserData(userData models.UserData) error
	UpdateUserData(userID int, userData models.UserData) error
	Close()
}
