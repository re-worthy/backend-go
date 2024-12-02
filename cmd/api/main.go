package main

import (
	"fmt"
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers"
	"github.com/re-worthy/backend-go/pkg/utils"
)

func main() {
	mux := http.NewServeMux()
	baseHandler := handlers.NewBaseHandler()
	handlers.SetupRoutes(mux, baseHandler)

	fmt.Println("Server starting on 0.0.0.0:8080...")
	errServer := http.ListenAndServe("0.0.0.0:8080", mux)
	utils.PanicOnError(&errServer)
}
