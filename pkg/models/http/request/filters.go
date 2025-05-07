package request

// Filters фильтр поиска
// @Description фильтр поиска людей
type Filters struct {
	Name       string `json:"name" form:"name" example:"Donald"`
	Surname    string `json:"surname" form:"surname" example:"Trump"`
	Patronymic string `json:"patronymic" form:"patronymic" example:"Duck"`
	Age        int    `json:"age" form:"age" example:"25"`
	Gender     string `json:"gender" form:"gender" example:"male"`
	Nation     string `json:"nation" form:"nation" example:"US"`
	Limit      int    `json:"limit" form:"limit" example:"3"`
	Offset     int    `json:"offset" form:"offset" example:"2"`
}
