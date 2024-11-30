package configs

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type Config struct {
	Env          string     `yaml:"env" env:"ENV" env-required:"true"`
	Storage_path string     `yaml:"storage_path" env-required:"true"`
	HttpServer   HttpServer `yaml:"http_server"`
}

func ShouldLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configFlag := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *configFlag
		if configPath == "" {
			log.Fatal("Config path is Not Set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config File Does Not Exist - %s/n", configPath)
	}

	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		fmt.Println("cannot read config file")
		log.Fatal("err: ", err)
	}
	return &config
}
