package utils

import (
	"service/pkg/models"
	"service/pkg/models/domain/filters"
	"service/pkg/models/domain/user"
	"service/pkg/models/http/request"
	"service/pkg/models/http/response"
)

// HttpDataToDomainFullUser - конвертирование http-модели в domain-модель
func HttpDataToDomainFullUser(reqUser request.User, userId models.UserId) *user.FullUser {
	resp := &user.FullUser{
		User: user.User{
			Name:       reqUser.Name,
			Surname:    reqUser.Surname,
			Patronymic: reqUser.Patronymic,
			Age:        reqUser.Age,
			Gender:     reqUser.Gender,
			Nation:     reqUser.Nation,
		},
		Id: userId,
	}
	return resp
}

// HttpDataToDomainFilters - конвертирование http-модели в domain-модель
func HttpDataToDomainFilters(reqFilters request.Filters) *filters.Filters {
	resp := &filters.Filters{
		Name:       reqFilters.Name,
		Surname:    reqFilters.Surname,
		Patronymic: reqFilters.Patronymic,
		Age:        reqFilters.Age,
		Gender:     reqFilters.Gender,
		Nation:     reqFilters.Nation,
		Limit:      reqFilters.Limit,
		Offset:     reqFilters.Offset,
	}
	return resp
}

// HttpDataToDomainCreateUser - конвертирование http-модели в domain-модель
func HttpDataToDomainCreateUser(reqUserInput request.UserInput) *user.UserInput {
	resp := &user.UserInput{
		Name:       reqUserInput.Name,
		Surname:    reqUserInput.Surname,
		Patronymic: reqUserInput.Patronymic,
	}
	return resp
}

// DomainDataToHttpFullUser - конвертирует domain-модель в объект http-ответа объекта FullUser
func DomainDataToHttpFullUser(domainFullUser user.FullUser) *response.FullUser {
	resp := &response.FullUser{
		User: response.User{
			Name:       domainFullUser.Name,
			Surname:    domainFullUser.Surname,
			Patronymic: domainFullUser.Patronymic,
			Age:        domainFullUser.Age,
			Gender:     domainFullUser.Gender,
			Nation:     domainFullUser.Nation,
		},
		Id: domainFullUser.Id,
	}
	return resp
}

func DomainDataToHttpFullUsers(domainFullUsers []user.FullUser) []response.FullUser {
	resp := make([]response.FullUser, len(domainFullUsers))
	for i, domainFullUser := range domainFullUsers {
		resp[i] = *DomainDataToHttpFullUser(domainFullUser)
	}
	return resp
}
