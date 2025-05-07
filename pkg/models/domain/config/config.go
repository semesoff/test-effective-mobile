package config

type Config struct {
	Database Database
	App      App
	Enrich   Enrich
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Driver   string
}

type App struct {
	Port string
}

type Enrich struct {
	UrlAge    string
	UrlGender string
	UrlNation string
}
