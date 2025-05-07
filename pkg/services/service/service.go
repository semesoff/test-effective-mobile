package service

import (
	"service/internal/db"
	"service/pkg/models"
	"service/pkg/models/domain/config"
	"service/pkg/models/domain/filters"
	"service/pkg/models/domain/user"
	"service/pkg/services/enrichment_service"
)

type ServiceManager struct {
	enrichmentService enrichment_service.EnrichmentService
	db                db.Database
}

type Service interface {
	CreateUser(user user.UserInput) (user.FullUser, error)
	ChangeUser(user user.FullUser) (user.FullUser, error)
	DeleteUser(userId models.UserId) error
	GetUsers(filters filters.Filters) ([]user.FullUser, error)
}

func NewServiceManager(config config.Database, configEnrich config.Enrich) *ServiceManager {
	serviceManager := &ServiceManager{}
	serviceManager.enrichmentService = enrichment_service.NewEnrichmentServiceManager(configEnrich)
	serviceManager.db = db.NewDatabaseManager(config)
	return serviceManager
}

func (s *ServiceManager) CreateUser(domainUser user.UserInput) (user.FullUser, error) {
	newUser := user.User{
		Name:       domainUser.Name,
		Surname:    domainUser.Surname,
		Patronymic: domainUser.Patronymic,
	}

	if err := s.enrichmentService.EnrichmentUser(&newUser); err != nil {
		return user.FullUser{}, err
	}

	fullUser, err := s.db.CreateUser(newUser)

	if err != nil {
		return user.FullUser{}, err
	}

	return fullUser, nil
}

func (s *ServiceManager) ChangeUser(domainUser user.FullUser) (user.FullUser, error) {
	changedUser, err := s.db.ChangeUser(domainUser)
	if err != nil {
		return user.FullUser{}, err
	}

	return changedUser, nil
}

func (s *ServiceManager) DeleteUser(userId models.UserId) error {
	err := s.db.DeleteUser(userId)
	return err
}

func (s *ServiceManager) GetUsers(filters filters.Filters) ([]user.FullUser, error) {
	users, err := s.db.GetUsers(filters)
	if err != nil {
		return []user.FullUser{}, err
	}
	return users, nil
}
