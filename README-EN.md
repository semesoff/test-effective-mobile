# Test task for Effective Mobile

The implementation of the service that will receive
a name by the API of the API from open APIs to enrich the answer with the most likely age,
floor and nationality and maintain data in the database. 
On request, issue information about the people found.

# Run service

## Clone repository
```bash
git clone https://github.com/semesoff/test-effective-mobile.git
```

## Change to project directory
```bash
cd test-effective-mobile
```

## Run containers
```bash
docker-compose up -d
```

## Init swagger (optional)
```bash
swag init -g ./pkg/routes/routes.go -o ./docs
```

# Swagger url
```bash
http://localhost:8080/swagger/index.html
```

# API Endpoints

## Get Users
`GET /api/users`

Used to get a list of users with filtering options.

### Query Parameters
- `name` - user's first name (string)
- `surname` - last name (string)
- `patronymic` - middle name (string)
- `age` - age (number)
- `gender` - gender (string)
- `nation` - nationality (string)
- `limit` - limit number of records (number)
- `offset` - offset (number)

### Request Example
```bash
GET /api/users?name=Donald&age=25&gender=male&nation=US&limit=3&offset=2
```

### Response
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

## Create User
`POST /api/users`

Create a new user with basic information.

### Request Body
```json
{
    "name": "Donald",
    "surname": "Trump",
    "patronymic": "Duck"
}
```

### Response (201 Created)
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

## Update User
`PUT /api/users/{id}`

Update existing user data.

### Path Parameters
- `id` - user identifier

### Request Body
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

### Response
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

## Delete User
`DELETE /api/users/{id}`

Delete user by identifier.

### Path Parameters
- `id` - user identifier

### Response
```json
{
    "message": "User deleted"
}
```