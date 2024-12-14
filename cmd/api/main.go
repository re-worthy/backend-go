package main

import (
	"log"
	"net/http"

	env_init "github.com/re-worthy/backend-go/internal/env"
	"github.com/re-worthy/backend-go/internal/handlers"
	"github.com/re-worthy/backend-go/pkg/utils"
)

func init() {
	env_init.LoadEnv()
	env_init.ValidateEnv()
}

func main() {
	mux := http.NewServeMux()
	baseHandler := handlers.NewBaseHandler()
	handlers.SetupRoutes(mux, baseHandler)

	log.Println("Server starting on 0.0.0.0:8080...")
	errServer := http.ListenAndServe("0.0.0.0:8080", mux)
	utils.PanicOnError(&errServer)
}
