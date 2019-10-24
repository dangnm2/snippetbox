package main

import (
	"os"
)

var (
	ADDRESS_ENV     = "SNIPPETBOX_ADDRESS"
	DB_USER_ENV     = "SNIPPETBOX_DB_USER"
	DB_PASSWORD_ENV = "SNIPPETBOX_DB_PASSWORD"
	DB_ADDRESS_ENV  = "SNIPPETBOX_DB_ADDRESS"
	DB_NAME_ENV     = "SNIPPETBOX_DB_NAME"
	SECRET_ENV      = "SNIPPETBOX_SECRET"
)

type DBConfig struct {
	User     string
	Password string
	Name     string
	Address  string
}

type Config struct {
	Address string
	DB      DBConfig
	Secret  string
}

func newConfig() Config {
	return Config{
		Address: getEnv(ADDRESS_ENV, "127.0.0.1:4000"),
		Secret:  getEnv(SECRET_ENV, "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"),
		DB: DBConfig{
			User:     getEnv(DB_USER_ENV, ""),
			Password: getEnv(DB_PASSWORD_ENV, ""),
			Name:     getEnv(DB_NAME_ENV, ""),
			Address:  getEnv(DB_ADDRESS_ENV, ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
