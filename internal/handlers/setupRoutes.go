package handlers

import (
	"net/http"

	helloworld "github.com/re-worthy/backend-go/internal/handlers/services/helloWorld"
	"github.com/re-worthy/backend-go/internal/handlers/services/users"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

func SetupRoutes(mux *http.ServeMux, h *handlers.TBaseHandler) {
	mux.HandleFunc("POST /", handlers.Adapter(helloworld.HelloWhoHandler, h))
	mux.HandleFunc("GET /", handlers.Adapter(helloworld.HelloWorldHandler, h))
	mux.HandleFunc("GET /inc", handlers.Adapter(helloworld.HelloDBHandler, h))
	mux.HandleFunc("POST /users", handlers.Adapter(users.CreateHandler, h))
	mux.HandleFunc("GET /users/{user_id}", handlers.Adapter(users.GetOneHandler, h))
}
