package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"service/pkg/models"
	"service/pkg/models/domain/config"
	"service/pkg/models/http/request"
	"service/pkg/services/service"
	"service/pkg/utils"
	"strconv"
)

type HandlersManager struct {
	service service.Service
}

type Handlers interface {
	GetUsers(c *gin.Context)
	DeleteUser(c *gin.Context)
	ChangeUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

func NewHandlersManager(config config.Database, configEnrich config.Enrich) *HandlersManager {
	handlersManager := &HandlersManager{}
	handlersManager.service = service.NewServiceManager(config, configEnrich)
	return handlersManager
}

// GetUsers - получение пользователей
// @Summary Получение пользователей
// @Description получение пользователей по фильтру
// @Tags users
// @Accept json
// @Produce json
// @Param filters query request.Filters false "Фильтры для получения пользователей"
// @Success 200 {object} []response.FullUser
// @Failure 400 {object} map[string]string "error": "Invalid format"
// @Failure 500 {object} map[string]string "error": "error message"
// @Router /users [get]
func (h *HandlersManager) GetUsers(c *gin.Context) {
	var filters request.Filters
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid format: %s", err.Error())})
		return
	}
	logrus.Debugf("GetUsers: Received filters from client: %+v", filters)
	domainFilters := utils.HttpDataToDomainFilters(filters)

	users, err := h.service.GetUsers(*domainFilters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseData := utils.DomainDataToHttpFullUsers(users)

	c.JSON(http.StatusOK, responseData)
}

// DeleteUser - удаление пользователя
// @Summary Удаление пользователя
// @Description удаление пользователя по id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]string "message": "User deleted"
// @Failure 400 {object} map[string]string "error": "Invalid format"
// @Failure 500 {object} map[string]string "error": "error message"
// @Router /users/{id} [delete]
func (h *HandlersManager) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil || userId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	if err := h.service.DeleteUser(models.UserId(userId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// ChangeUser - изменение данных пользователя
// @Summary Изменение данных пользователя
// @Description Изменение данных пользователя по некоторым данным
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param newUserData body request.User true "User"
// @Success 200 {object} models.FullUser
// @Failure 400 {object} map[string]string "error": "Invalid format"
// @Failure 500 {object} map[string]string "error": "error message"
// @Router /users/{id} [put]
func (h *HandlersManager) ChangeUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil || userId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	var user request.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	fullUser := utils.HttpDataToDomainFullUser(user, models.UserId(userId))

	changedUser, err := h.service.ChangeUser(*fullUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseData := utils.DomainDataToHttpFullUser(changedUser)

	c.JSON(http.StatusOK, responseData)
}

// CreateUser - создание пользователя
// @Summary Создание пользователя
// @Description Создание нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.UserInput true "User Input"
// @Success 201 {object} response.FullUser
// @Failure 400 {object} map[string]string "error": "Invalid format"
// @Failure 500 {object} map[string]string "error": "error message"
// @Router /users [post]
func (h *HandlersManager) CreateUser(c *gin.Context) {
	var user request.UserInput
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
		return
	}

	domainCreateData := utils.HttpDataToDomainCreateUser(user)

	fullUser, err := h.service.CreateUser(*domainCreateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseData := utils.DomainDataToHttpFullUser(fullUser)

	c.JSON(http.StatusCreated, responseData)
}
