package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	API struct {
		Host string `envconfig:"HOST" default:":5050"`
	}
	DB struct {
		Driver string `envconfig:"DRIVER" required:"true"`
		Mysql  struct {
			URL string `envconfig:"URL" required:"true"`
		}
	}
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfig(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("CATALOG", config); err != nil {
		return nil, err
	}
	return config, nil
}
