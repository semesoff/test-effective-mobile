package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"service/pkg/models"
	"service/pkg/models/domain/config"
	"service/pkg/models/domain/filters"
	"service/pkg/models/domain/user"
)

type DatabaseManager struct {
	db *sql.DB
}

type Database interface {
	CreateUser(user user.User) (user.FullUser, error)
	ChangeUser(user user.FullUser) (user.FullUser, error)
	DeleteUser(userId models.UserId) error
	GetUsers(filters filters.Filters) ([]user.FullUser, error)
}

func NewDatabaseManager(config config.Database) *DatabaseManager {
	databaseManager := &DatabaseManager{}
	databaseManager.Init(config)
	return databaseManager
}

func (dm *DatabaseManager) Init(config config.Database) {
	logrus.Debug("Initializing database...")
	db, err := sql.Open(config.Driver,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.Database,
			"disable",
		),
	)
	if err != nil {
		logrus.Fatal("Error initializing database: ", err)
		return
	}

	if err := db.Ping(); err != nil {
		logrus.Fatal("Error pinging database: ", err)
		return
	}

	dm.db = db

	logrus.Info("Database is initialized")
}
