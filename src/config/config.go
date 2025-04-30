package config

import (
	"log"
	"os"

	viper "github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	Cors     CorsConfig     `yaml:"cors"`
	Postgres PostgresConfig `yaml:"postgres"`
	Redis    RedisConfig    `yaml:"redis"`
}

type ServerConfig struct {
	Port    string `yaml:"Port"`
	RunMode string `yaml:"RunMode"`
}

type LoggerConfig struct {
	FilePath string `yaml:"FilePath"`
	Encoding string `yaml:"encoding"`
	Level    string `yaml:"level"`
}

type CorsConfig struct {
	AllowOrigins  []string `yaml:"allowOrigins"`
	AllowMethods  []string `yaml:"allowMethods"`
	AllowHeaders  []string `yaml:"allowHeaders"`
	ExposeHeaders []string `yaml:"exposeHeaders"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Db          int    `yaml:"db"`
	Password    string `yaml:"password"`
	PoolSize    int    `yaml:"poolSize"`
	PoolTimeout int    `yaml:"poolTimeout"`
}

func GetConfig(env string) (*Config, error) {
	path := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(path, "yaml")
	if err != nil {
		log.Printf("Error reading config file, %v", err)
		return nil, err
	}
	config, err := ParseConfig(v)
	if err != nil {
		log.Printf("Error parsing config file, %v", err)
		return nil, err
	}
	return config, nil

}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	var err error
	v.Unmarshal(&config)
	return &config, err
}

func LoadConfig(path string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigName(path)
	v.AddConfigPath("../config")
	err := v.ReadInConfig()
	return v, err
}

func getConfigPath(env string) string {
	if env == "dev" {
		return "config-dev.yml"
	} else if env == "prod" {
		return "config-prod.yml"
	}
	return "config-dev.yml"
}
