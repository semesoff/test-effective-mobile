package request

// UserInput Входные данные для создания пользователя
// @Description Входные данные для создания пользователя
type UserInput struct {
	Name       string `json:"name" binding:"required" example:"Donald"`
	Surname    string `json:"surname" binding:"required" example:"Trump"`
	Patronymic string `json:"patronymic" example:"Duck"`
}

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
