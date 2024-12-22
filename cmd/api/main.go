package main

import (
	"fmt"
	"log"
	"net/http"

	env_init "github.com/re-worthy/backend-go/internal/env"
	"github.com/re-worthy/backend-go/internal/handlers"
	"github.com/re-worthy/backend-go/pkg/utils"
)

var envConfig env_init.TEnvConfig

func init() {
	env_init.LoadEnv()
	envConfig = env_init.ValidateEnv()
}

func main() {
	mux := http.NewServeMux()
	baseHandler, dbOnClose, err := handlers.NewBaseHandler(envConfig)
	if err != nil {
		log.Fatalf("Error creating new base handler:\n\t%s", err.Error())
	}
	defer dbOnClose()

	handlers.SetupRoutes(mux, baseHandler)

	address := fmt.Sprintf("%s:%d", envConfig.SELF_HOST, envConfig.SELF_PORT)
	log.Printf("Server starting on %s ...", address)
	errServer := http.ListenAndServe(address, mux)
	utils.PanicOnError(&errServer)
}
