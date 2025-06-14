package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"telegram_gateway_service/pkg/logger"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Env string

var (
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"
)

type Config struct {
	*jsonConfig
	*envConfig
}

type jsonConfig struct {
	Server  ServerConfig  `json:"server"`
	Handler HandlerConfig `json:"handler"`
}

type envConfig struct {
	Env Env `env:"ENV" env-required:"true"`
}

type ServerConfig struct {
	Port           int           `json:"port"`
	ReadTimeout    time.Duration `json:"read_timeout"`
	WriteTimeout   time.Duration `json:"write_timeout"`
	MaxHeaderBytes int           `json:"max_header_bytes"`
}

type HandlerConfig struct {
	RequestTimeout  time.Duration `json:"request_timeout"`
	RegisterTimeout time.Duration `json:"register_timeout"`
}

func MustConfig(log logger.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		log.Panic(fmt.Sprintf("failed download the file .env: %v", err))
	}

	path := fetchConfigPath()
	if path == "" {
		log.Panic("config path is empty")
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Panic("config file does not exist: " + path)
	}
	viper.AddConfigPath(filepath.Dir(path))
	viper.SetConfigType("json")
	viper.SetConfigName(filepath.Base(path))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	var jsonCfg jsonConfig

	err = viper.Unmarshal(&jsonCfg)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}
	validate := validator.New()

	err = validate.Struct(jsonCfg)
	if err != nil {
		log.Panicf("unable to validate config file: %v", err)
	}
	var envCfg envConfig

	err = cleanenv.ReadEnv(&envCfg)
	if err != nil {
		log.Panic("failed to read envConfig: " + err.Error())
	}
	return &Config{
		jsonConfig: &jsonCfg,
		envConfig:  &envCfg,
	}
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
