package routes

// @title API Service
// @version 1.0.0
// @description Enrichment Service
// @host localhost:8080
// @BasePath /api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"service/pkg/handlers"
	_ "service/pkg/models"
)

type RoutesManager struct {
	handlers handlers.Handlers
}

type Routes interface {
	Init(*gin.Engine)
}

func NewRoutesManager(handlers handlers.Handlers) *RoutesManager {
	return &RoutesManager{
		handlers: handlers,
	}
}

func (rm *RoutesManager) Init(engine *gin.Engine) {
	// Swagger endpoint
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("api")
	api.POST("/users", rm.handlers.CreateUser)
	api.GET("/users", rm.handlers.GetUsers)
	api.DELETE("/users/:id", rm.handlers.DeleteUser)
	api.PUT("/users/:id", rm.handlers.ChangeUser)
}
