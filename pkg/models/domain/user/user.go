package user

import (
	"service/pkg/models"
)

// UserInput - вводные данные
type UserInput struct {
	Name       string
	Surname    string
	Patronymic string
}

// User Данные пользователя с дополнительной информацией
type User struct {
	Name       string
	Surname    string
	Patronymic string
	Age        int
	Gender     string
	Nation     string
}

// FullUser - полные данные пользователя
type FullUser struct {
	User
	Id models.UserId
}
