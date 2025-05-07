package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"service/pkg/models"
	"service/pkg/models/domain/filters"
	"service/pkg/models/domain/user"
	"strings"
)

// CreateUser - создание нового пользователя
func (dm *DatabaseManager) CreateUser(domainUser user.User) (user.FullUser, error) {
	var newUser user.FullUser
	err := dm.db.QueryRow(
		"INSERT INTO users (name, surname, patronymic, age, gender, nation) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		domainUser.Name,
		domainUser.Surname,
		domainUser.Patronymic,
		domainUser.Age,
		domainUser.Gender,
		domainUser.Nation,
	).Scan(&newUser.Id)
	if err != nil {
		return user.FullUser{}, err
	}
	newUser.User = domainUser
	return newUser, nil
}

// ChangeUser - изменение данных пользователя
func (dm *DatabaseManager) ChangeUser(domainUser user.FullUser) (user.FullUser, error) {
	var exists bool
	err := dm.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", domainUser.Id).Scan(&exists)
	if err != nil {
		return user.FullUser{}, err
	}
	if !exists {
		return user.FullUser{}, fmt.Errorf("domainUser with id %d not found", domainUser.Id)
	}

	_, err = dm.db.Exec("UPDATE users SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nation = $6 WHERE id = $7",
		domainUser.Name,
		domainUser.Surname,
		domainUser.Patronymic,
		domainUser.Age,
		domainUser.Gender,
		domainUser.Nation,
		domainUser.Id,
	)
	if err != nil {
		return user.FullUser{}, err
	}
	return domainUser, nil
}

// DeleteUser - удаление пользователя
func (dm *DatabaseManager) DeleteUser(userId models.UserId) error {
	var exists bool
	err := dm.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)",
		userId,
	).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user with id %d not found", userId)
	}

	_, err = dm.db.Exec("DELETE FROM users WHERE id = $1", userId)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers - получение списка пользователей по фильтрации
func (dm *DatabaseManager) GetUsers(filters filters.Filters) ([]user.FullUser, error) {
	query, args := dm.createQuery(filters)
	logrus.Debug("Query for GetUsers: ", query)
	rows, err := dm.db.Query(query, args...)
	if err != nil {
		return []user.FullUser{}, err
	}

	users := make([]user.FullUser, 0)
	for rows.Next() {
		var fullUser user.FullUser
		if err := rows.Scan(
			&fullUser.Id,
			&fullUser.Name,
			&fullUser.Surname,
			&fullUser.Patronymic,
			&fullUser.Age,
			&fullUser.Gender,
			&fullUser.Nation,
		); err != nil {
			return []user.FullUser{}, err
		}
		users = append(users, fullUser)
	}
	return users, nil
}

// createQuery - создает запрос фильтрации для метода GetUsers
func (dm *DatabaseManager) createQuery(filters filters.Filters) (string, []interface{}) {
	logrus.Debug("Filters for GetUsers(createQuery): ", filters)
	var query strings.Builder
	query.WriteString("SELECT * FROM users")

	var conditions []string
	var args []interface{}

	val := reflect.ValueOf(filters)
	typ := val.Type()

	k := 1
	for i := 0; i < val.NumField(); i++ {
		// получение значения i-ого поля
		field := val.Field(i)
		if !field.IsValid() || field.IsZero() {
			continue
		}
		// получение имени поля
		colName := strings.ToLower(typ.Field(i).Name)

		// пропускаем LIMIT и OFFSET
		if colName == "limit" || colName == "offset" {
			continue
		}

		// получение значения из поля
		val := field.Interface()

		// запись условий и аргументов
		conditions = append(conditions, fmt.Sprintf("%s = $%d", colName, k))
		args = append(args, val)
		k++
	}

	// формирование полного условия WHEN
	if len(conditions) > 0 {
		query.WriteString(" WHERE " + strings.Join(conditions, " AND "))
	}

	// формирование окончательного запроса
	if filters.Limit > 0 {
		query.WriteString(fmt.Sprintf(" LIMIT %d", filters.Limit))
	}
	if filters.Offset > 0 {
		query.WriteString(fmt.Sprintf(" OFFSET %d", (filters.Offset-1)*filters.Limit))
	}

	return query.String(), args
}
