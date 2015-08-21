package engine

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port   int    `envconfig:"PORT"`
	DBHost string `envconfig:"DB_HOST"`
	DBName string `envconfig:"DB_NAME"`
}

var env Environment

func LoadEnv() {
	env.Port = 8080
	env.DBHost = "localhost:27017"
	env.DBName = "miri-dev"

	err := envconfig.Process("MIRI", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Info("\nEnvironment\n-----------\nPort: %v\nDB Host: %s\nDB Name: %s\n",
		env.Port, env.DBHost, env.DBName)
}
