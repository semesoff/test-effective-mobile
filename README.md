# Тестовое задание для Effective Mobile

[English version](README-EN.md)

Реализация сервиса, который будет получать по API ФИО, из открытых API обогащать
ответ наиболее вероятными возрастом, полом и национальностью и сохранять данные в
БД. По запросу выдавать информацию о найденных людях.

# Запуск сервиса

## Клонирование репозитория
```bash
git clone https://github.com/semesoff/test-effective-mobile.git
```

## Переход в директорию проекта
```bash
cd test-effective-mobile
```

## Запуск контейнеров
```bash
docker-compose up -d
```

## Инициализация swagger (необязательно)
```bash
swag init -g ./pkg/routes/routes.go -o ./docs
```

# Адрес swagger
```bash
http://localhost:8080/swagger/index.html
```

# API Endpoints

## Получение пользователей
`GET /api/users`

Используется для получения списка пользователей с возможностью фильтрации.

### Параметры запроса
- `name` - имя пользователя (строка)
- `surname` - фамилия (строка)
- `patronymic` - отчество (строка)
- `age` - возраст (число)
- `gender` - пол (строка)
- `nation` - национальность (строка)
- `limit` - ограничение количества записей (число)
- `offset` - смещение (число)

### Пример запроса
```bash
GET /api/users?name=Donald&age=25&gender=male&nation=US&limit=3&offset=2
```

### Ответ
```json
[
    {
        "id": 1,
        "name": "Donald",
        "surname": "Trump",
        "patronymic": "Duck",
        "age": 25,
        "gender": "male",
        "nation": "US"
    }
]
```

## Создание пользователя
`POST /api/users`

Создание нового пользователя с базовой информацией.

### Тело запроса
```json
{
    "name": "Donald",
    "surname": "Trump",
    "patronymic": "Duck"
}
```

### Ответ (201 Created)
```json
{
    "id": 1,
    "name": "Donald",
    "surname": "Trump",
    "patronymic": "Duck",
    "age": 25,
    "gender": "male",
    "nation": "US"
}
```

## Изменение пользователя
`PUT /api/users/{id}`

Обновление данных существующего пользователя.

### Параметры пути
- `id` - идентификатор пользователя

### Тело запроса
```json
{
    "name": "Donald",
    "surname": "Trump",
    "patronymic": "Duck",
    "age": 25,
    "gender": "male",
    "nation": "US"
}
```

### Ответ
```json
{
    "id": 1,
    "name": "Donald",
    "surname": "Trump",
    "patronymic": "Duck",
    "age": 25,
    "gender": "male",
    "nation": "US"
}
```

## Удаление пользователя
`DELETE /api/users/{id}`

Удаление пользователя по идентификатору.

### Параметры пути
- `id` - идентификатор пользователя

### Ответ
```json
{
    "message": "User deleted"
}
```