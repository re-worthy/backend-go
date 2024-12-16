package env_init

import (
	"log"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type TEnvConfig struct {
	DATABASE_URL        string `env:"DATABASE_URL,notEmpty,required"`
	DATABASE_AUTH_TOKEN string `env:"DATABASE_AUTH_TOKEN,required"`
	SELF_HOST           string `env:"SELF_HOST,notEmpty,required"`
	ENVIRONMENT         string `env:"ENVIRONMENT,notEmpty,required"`
	SELF_PORT           int    `env:"SELF_PORT,notEmpty,required"`
}

func ValidateEnv() TEnvConfig {
	var cfg TEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Error validating environment variables:\n\t%s", err)
	}

	log.Println("Validated environment")
	return cfg
}

func LoadEnv() {
	const (
		ENV_VAR_NAME        = "ENVIRONMENT"
		ENV_NAME_PRODUCTION = "PRODUCTION"
		ENV_NAME_DEV        = "DEVELOPMENT"
		ENV_DEFAULT         = ENV_NAME_PRODUCTION
		ENV_FILENAME_PROD   = ".env"
		ENV_FILENAME_DEV    = ".env.dev"
	)

	environment := os.Getenv(ENV_VAR_NAME)
	if environment == "" {
		log.Println("ENVIRONMENT variable not found, defaulting to \"production\"")
		environment = ENV_DEFAULT
	}

	if environment == ENV_NAME_PRODUCTION {
		err := godotenv.Load(ENV_FILENAME_PROD)
		if err != nil {
			log.Fatalf("Error loading %s file:\n\t%s", ENV_FILENAME_PROD, err)
		}
		log.Printf("Loaded environment variables for env:%s", environment)
	}

	if environment == ENV_NAME_DEV {
		errDev := godotenv.Load(ENV_FILENAME_DEV)
		if errDev != nil {
			log.Fatalf("Error loading %s file:\n\t%s", ENV_FILENAME_DEV, errDev)
		}
		log.Printf("Loaded environment variables for env:%s", environment)
	}
}
