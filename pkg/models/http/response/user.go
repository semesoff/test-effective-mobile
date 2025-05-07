package response

import (
	"service/pkg/models"
)

// User Данные пользователя с дополнительной информацией
// @Description Данные пользователя с дополнительной информацией
type User struct {
	Name       string `json:"name" example:"Donald"`
	Surname    string `json:"surname" example:"Trump"`
	Patronymic string `json:"patronymic" example:"Duck"`
	Age        int    `json:"age" example:"25"`
	Gender     string `json:"gender" example:"male"`
	Nation     string `json:"nation" example:"US"`
}

// FullUser представляет полные данные пользователя
// @Description Полные данные пользователя с id
type FullUser struct {
	User
	Id models.UserId `json:"id" example:"1"`
}
