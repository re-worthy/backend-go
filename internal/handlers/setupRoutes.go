package handlers

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/services"
)

func SetupRoutes(mux *http.ServeMux, h *BaseHandler) {
	mux.HandleFunc("POST /", Adapter(services.HelloWorldService.HelloWhoHandler))
	mux.HandleFunc("GET /", Adapter(services.HelloWorldService.HelloWorldHandler))
}
