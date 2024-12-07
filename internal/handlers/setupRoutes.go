package handlers

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/services/helloWorld"
	"github.com/re-worthy/backend-go/internal/handlers/types"
)

func SetupRoutes(mux *http.ServeMux, h *BaseHandler) {
	mux.HandleFunc("POST /", handlers.Adapter(helloworld.HelloWhoHandler))
	mux.HandleFunc("GET /", handlers.Adapter(helloworld.HelloWorldHandler))
}
