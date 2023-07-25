package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Database string `mapstructure:"db"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"pass"`
	Database int    `mapstructure:"db"`
}

type MailchimpHTTPConfig struct {
	BaseURL string `mapstructure:"base"`
	APIKey  string `mapstructure:"apikey"`
	Server  string `mapstructure:"server"`
}

type OmetriaHTTPConfig struct {
	BaseURL string `mapstructure:"base"`
	APIKey  string `mapstructure:"apikey"`
}

type HTTPConfig struct {
	Mailchimp MailchimpHTTPConfig `mapstructure:"mailchimp"`
	Ometria   OmetriaHTTPConfig   `mapstructure:"ometria"`
}

type AppConfig struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	HTTP     HTTPConfig
}

func LoadConfigs(path, filename string) AppConfig {
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("cannot read config file %s from path %s", filename, path))
	}
	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("cannot unmarshal file %s from path %s", filename, path))
	}

	config.HTTP.Ometria.APIKey = os.Getenv("OMETRIA_APIKEY")
	config.HTTP.Mailchimp.APIKey = os.Getenv("MAILCHIMP_APIKEY")

	return config
}

func NewHTTPClient() http.Client {
	return http.Client{
		Timeout: time.Second * 60,
	}
}
