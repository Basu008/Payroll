package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

// This file contains all the configuration that will be needed in our project
type Config struct {
	ServerConfig    ServerConfig    `mapstructure:"server"`
	APIConfig       APIConfig       `mapstructure:"api"`
	APPConfig       APPConfig       `mapstructure:"app"`
	DatabaseConfig  DatabaseConfig  `mapstructure:"database"`
	TokenAuthConfig TokenAuthConfig `mapstructure:"token"`
}

// Every config that is needed to run our server
type ServerConfig struct {
	ListenAddr         string             `mapstructure:"listenAddr"`
	Port               string             `mapstructure:"port"`
	ReadTimeout        time.Duration      `mapstructure:"readTimeout"`
	WriteTimeout       time.Duration      `mapstructure:"writeTimeout"`
	CloseTimeout       time.Duration      `mapstructure:"closeTimeout"`
	ApiEndPointsConfig ApiEndPointsConfig `mapstructure:"apiEndPointsConfig"`
	CORSConfig         CORSConfig         `mapstructure:"cors"`
	Env                string             `mapstructure:"env"`
}

type ApiEndPointsConfig struct {
	BaseAPI string `mapstructure:"baseApi"`
}

type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowedOrigins"`
	AllowedMethods   []string `mapstructure:"allowedMethods"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
	AllowedHeaders   []string `mapstructure:"allowedHeaders"`
}

type APIConfig struct {
	Mode               string `mapstructure:"mode"`
	EnableTestRoute    bool   `mapstructure:"enableTestRoute"`
	MaxRequestDataSize int    `mapstructure:"maxRequestDataSize"`
	ApiEndPointsConfig ApiEndPointsConfig
}

type APPConfig struct {
	DatabaseConfig     DatabaseConfig
	ApiEndPointsConfig ApiEndPointsConfig
	EmployeeConfig     ServiceConfig `mapstructure:"employee"`
}

type DatabaseConfig struct {
	Scheme     string `mapstructure:"scheme"`
	Host       string `mapstructure:"host"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	ReplicaSet string `mapstructure:"replicaSet"`
}

func (d *DatabaseConfig) GetConnectionURL() string {
	url := fmt.Sprintf("%s://", d.Scheme)
	if d.Username != "" && d.Password != "" {
		url += fmt.Sprintf("%s:%s@", d.Username, d.Password)
	}
	url += d.Host
	return url
}

type ServiceConfig struct {
	DBName string `mapstructure:"dbName"`
}

type TokenAuthConfig struct {
	JWTSignKey string `mapstructure:"jwtSignKey"`
}

func GetConfig(filename string) *Config {
	viper.SetConfigName(filename)
	viper.AddConfigPath("../conf/")
	viper.AddConfigPath("../../conf/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load config: %s", err)
		os.Exit(1)
	}
	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	config.APIConfig.ApiEndPointsConfig = config.ServerConfig.ApiEndPointsConfig
	config.APPConfig.ApiEndPointsConfig = config.ServerConfig.ApiEndPointsConfig
	return config
}
