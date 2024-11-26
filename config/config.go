package config

import (
	"maps-house/pkg/logger"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress string
	OsuApiKey     string
	SecretKey     string
	Workers       int
	MapsPath      string
	DB_Host       string `mapstructure:"POSTGRES_HOST"`
	DB_Name       string `mapstructure:"POSTGRES_DB"`
	DB_Port       string `mapstructure:"POSTGRES_PORT"`
	DB_User       string `mapstructure:"POSTGRES_USER"`
	DB_Password   string `mapstructure:"POSTGRES_PASSWORD"`
}

func NewConfig(logger *logger.Logger) (*Config, error) {
	c := Config{
		ListenAddress: ":8000",
		Workers:       2,
		MapsPath:      "/mnt/maps/maps",
	}
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	viper.SafeWriteConfig()
	return &c, nil
}

func (c *Config) DSNBuilder() string {
	var dsn strings.Builder
	dsn.WriteString("host=")
	dsn.WriteString(c.DB_Host)
	dsn.WriteString(" user=")
	dsn.WriteString(c.DB_User)
	dsn.WriteString(" password=")
	dsn.WriteString(c.DB_Password)
	dsn.WriteString(" dbname=")
	dsn.WriteString(c.DB_Name)
	dsn.WriteString(" port=")
	dsn.WriteString(c.DB_Port)

	dsn.WriteString(" sslmode=disable TimeZone=Europe/Moscow")

	return dsn.String()
}

func (c *Config) DSNMySQLBuilder() string {
	var dsn strings.Builder

	dsn.WriteString(c.DB_User)
	dsn.WriteString(":")
	dsn.WriteString(c.DB_Password)
	dsn.WriteString("@(")
	dsn.WriteString(c.DB_Host)
	dsn.WriteString(":")
	dsn.WriteString(c.DB_Port)
	dsn.WriteString(")/")
	dsn.WriteString(c.DB_Name)

	return dsn.String()
}
