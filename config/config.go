package config

import (
	"errors"
	"fmt"
	"io"
	"mapsHouse/pkg/logger"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DSN       string
	Host      string
	SecretKey string
	Workers   int
	DB        db
}

type db struct {
	Host     string
	DBname   string
	Port     string
	User     string
	Password string
}

var confFile = "./config.yml"

func NewConfig(logger *logger.Logger) (*Config, error) {
	c := Config{
		Workers: 2,
		DB: db{
			Host:     "127.0.0.1",
			DBname:   "ripple",
			Port:     "3306",
			User:     "",
			Password: "",
		},
	}

	f, err := os.Open(confFile)
	switch {
	case errors.Is(err, os.ErrNotExist):
		{
			logger.Warn("No config.yml was found. Creating config.yml file...")

			f, err = os.Create(confFile)
			if errors.Is(err, os.ErrPermission) {
				logger.Error("Can't create config.yml file, permission denied")
				return nil, err
			}
			if err != nil {
				logger.Error("Can't create config.yml file")
				return nil, err
			}

			b, err := yaml.Marshal(&c)
			if err != nil {
				logger.Error("Can't process config.yml")
				return nil, err
			}

			_, err = f.Write(b)
			if err != nil {
				logger.Error("Can't write bytes into config.yml")
				return nil, err
			}
		}
	case err != nil:
		{
			logger.Error(fmt.Sprintf("config.yml error: %v", err))
			return nil, err
		}
	}

	b, err := io.ReadAll(f)
	if err != nil {
		logger.Error("Can't read config.yml")
		return nil, err
	}

	yaml.Unmarshal(b, &c)
	return &c, nil
}

func (c *Config) DSNBuilder() string {
	var dsn strings.Builder

	dsn.WriteString(c.DB.User)
	dsn.WriteString(":")
	dsn.WriteString(c.DB.Password)
	dsn.WriteString("@(")
	dsn.WriteString(c.DB.Host)
	dsn.WriteString(":")
	dsn.WriteString(c.DB.Port)
	dsn.WriteString(")/")
	dsn.WriteString(c.DB.DBname)

	return dsn.String()
}
