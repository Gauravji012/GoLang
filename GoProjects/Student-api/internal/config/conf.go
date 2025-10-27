// it contains the internal packages of the project
// it is same as .pkg folder content
// hmne struct ko serialize kiya h uske liye ek go lang ka package use kiya h -> golang clean env package "go get -u github.com/ilyakaznacheev/cleanenv"

package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// we serialze the struct using “ backtick
// yaml:"env" -> this filed is present with env name in yaml file
type Config struct {
	Env         string               `yaml:"env" env:"ENV" env-required:"true"` // env:"ENV" -> take value from environment env-required: "true" -> it is a required field, env-default:"production" -> also use to take the default value in env variable
	StoragePath string               `yaml:"storage_path" env-required:"true"`  // -> called as struct text
	HTTPServer  `yaml:"http_server"` // embaded httpserver struct
}

// write parsing logic
// this function should run if not run then we should retun from here instead of returning some error that is why we use MustLoad name for easy understanding
func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	// if execution reaches at this line -> we have confiuragtion path value so now we set into struct

	var cfg Config // we serialize into Config struct

	err := cleanenv.ReadConfig(configPath, &cfg) // seraization is ddone using this line
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
