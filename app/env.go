package app

import (
	"log"

	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port               int    `envconfig:"PORT"`
	DBHost             string `envconfig:"DB_HOST"`
	DBName             string `envconfig:"DB_NAME"`
	JWTExpirationDelta int    `envconfig:"JWT_EXPIRY_DELTA"`
	JWTSecretKey       string `envconfig:"JWT_SECRET_KEY"`
}

var env Environment

func init() {
	env.Port = 8080
	env.DBHost = "127.0.0.1:27017"
	env.DBName = "miri-dev"
	env.JWTExpirationDelta = 72
	env.JWTSecretKey = "i23k8jnTghdfadGGrt32hgSGH42zSD53HaraaR48990A"

	err := envconfig.Process("MIRI", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger.Write.Info("\nEnvironment\n-----------\nPort: %v\nDB Host: %s\nDB Name: %s\n",
		env.Port, env.DBHost, env.DBName)
}
