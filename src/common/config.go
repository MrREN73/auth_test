package common

import (
	"fmt"
)

// Config describes configuration for db.
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
	SSLMode  string `json:"ssl"`
	Schema   string `json:"schema"`
	Debug    bool   `json:"debug"`
}

func (c Config) connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		c.Host, c.Port, c.User, c.Password, c.DB, c.SSLMode, c.Schema)
}
