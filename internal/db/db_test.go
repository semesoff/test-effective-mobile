package db

import (
	"service/pkg/models"
	"service/pkg/models/domain"
	"testing"
)

var config = domain.Database{
	Host:     "localhost",
	Port:     "5050",
	User:     "root",
	Password: "root",
	Database: "db",
	Driver:   "postgres",
}

func setupTestDB(t *testing.T) Database {
	db := NewDatabaseManager(config)
	return db
}

func TestDatabaseManager_CreateUser(t *testing.T) {
	db := setupTestDB(t)

	tests := []struct {
		name    string
		user    user.User
		wantErr bool
	}{
		{
			name: "Успешное создание пользователя",
			user: user.User{
				UserInput: user.UserInput{
					Name:    "Иван",
					Surname: "Иванов",
				},
				Age:    25,
				Gender: "male",
				Nation: "russian",
			},
			wantErr: false,
		},
		{
			name: "Отрицательный возраст",
			user: user.User{
				UserInput: user.UserInput{
					Name:    "Петр",
					Surname: "Петров",
				},
				Age:    -1,
				Gender: "male",
				Nation: "russian",
			},
			wantErr: true,
		},
		{
			name: "Пустое имя",
			user: user.User{
				UserInput: user.UserInput{
					Name:    "",
					Surname: "Сидоров",
				},
				Age:    30,
				Gender: "male",
				Nation: "russian",
			},
			wantErr: true,
		},
		{
			name: "Некорректный пол",
			user: user.User{
				UserInput: user.UserInput{
					Name:    "Анна",
					Surname: "Петрова",
				},
				Age:    25,
				Gender: "invalid",
				Nation: "russian",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.CreateUser(tt.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDatabaseManager_ChangeUser(t *testing.T) {
	db := setupTestDB(t)

	fullUser, _ := db.CreateUser(user.User{
		UserInput: user.UserInput{
			Name:    "Антон",
			Surname: "Антонов",
		},
		Age:    30,
		Gender: "male",
		Nation: "russian",
	})

	tests := []struct {
		name    string
		user    user.FullUser
		wantErr bool
	}{
		{
			name: "Успешное изменение пользователя",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Антон",
						Surname: "Антонов",
					},
					Age:    100,
					Gender: "male",
					Nation: "american",
				},
				Id: fullUser.Id,
			},
			wantErr: false,
		},
		{
			name: "Отрицательный возраст",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Антон",
						Surname: "Антонов",
					},
					Age:    -100,
					Gender: "male",
					Nation: "american",
				},
				Id: fullUser.Id,
			},
			wantErr: true,
		},
		{
			name: "Большой возраст",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Антон",
						Surname: "Антонов",
					},
					Age:    150,
					Gender: "male",
					Nation: "american",
				},
				Id: fullUser.Id,
			},
			wantErr: true,
		},
		{
			name: "Неверный пол",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Антон",
						Surname: "Антонов",
					},
					Age:    100,
					Gender: "null",
					Nation: "american",
				},
				Id: fullUser.Id,
			},
			wantErr: true,
		},
		{
			name: "Нет имени",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "",
						Surname: "Антонов",
					},
					Age:    100,
					Gender: "male",
					Nation: "american",
				},
				Id: fullUser.Id,
			},
			wantErr: true,
		},
		{
			name: "Неверный ID",
			user: user.FullUser{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Антон",
						Surname: "Антонов",
					},
					Age:    100,
					Gender: "male",
					Nation: "american",
				},
				Id: -100,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.ChangeUser(tt.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseManager_DeleteUser(t *testing.T) {
	db := setupTestDB(t)

	fullUser, _ := db.CreateUser(user.User{
		UserInput: user.UserInput{
			Name:    "Антон",
			Surname: "Антонов",
		},
		Age:    30,
		Gender: "male",
		Nation: "russian",
	})

	tests := []struct {
		name    string
		userId  user.UserId
		wantErr bool
	}{
		{
			name:    "Успешное удаление",
			userId:  user.UserId(fullUser.Id),
			wantErr: false,
		},
		{
			name:    "Отрицательный ID",
			userId:  user.UserId(-1),
			wantErr: false,
		},
		{
			name:    "Неизвестный ID",
			userId:  user.UserId(10 << 20),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := db.DeleteUser(tt.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseManager_GetUsers(t *testing.T) {
	db := setupTestDB(t)

	fullUser1, _ := db.CreateUser(user.User{
		UserInput: user.UserInput{
			Name:    "Donald",
			Surname: "Trump",
		},
		Age:    30,
		Gender: "male",
		Nation: "american",
	})

	fullUser2, _ := db.CreateUser(user.User{
		UserInput: user.UserInput{
			Name:    "April",
			Surname: "Snow",
		},
		Age:    55,
		Gender: "male",
		Nation: "russian",
	})

	tests := []struct {
		name    string
		filters user.Filters
		usersId map[int]bool
		wantErr bool
	}{
		{
			name: "Фильтр по имени и фамилии",
			filters: user.Filters{
				User: user.User{
					UserInput: user.UserInput{
						Name:    "Donald",
						Surname: "Trump",
					},
				},
			},
			usersId: map[int]bool{fullUser1.Id: true, fullUser2.Id: true},
			wantErr: false,
		},
		{
			name: "Фильтр по имени, возрасту, гендеру",
			filters: user.Filters{
				User: user.User{
					UserInput: user.UserInput{
						Name: "April",
					},
					Age:    55,
					Gender: "male",
				},
			},
			usersId: map[int]bool{fullUser1.Id: true, fullUser2.Id: true},
			wantErr: false,
		},
		{
			name: "Фильтр по национальности",
			filters: user.Filters{
				User: user.User{
					Nation: "russian",
				},
			},
			usersId: map[int]bool{fullUser1.Id: true, fullUser2.Id: true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := db.GetUsers(tt.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
			exist := false
			for _, user := range users {
				if _, ok := tt.usersId[user.Id]; ok {
					exist = true
					break
				}
			}
			if !exist {
				t.Errorf("GetUsers() added user before not found with this filters")
			}
		})
	}
}
