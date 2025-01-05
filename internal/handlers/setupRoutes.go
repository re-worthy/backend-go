package handlers

import (
	"net/http"

	authService "github.com/re-worthy/backend-go/internal/handlers/services/auth"
	transactionsService "github.com/re-worthy/backend-go/internal/handlers/services/transactions"
	usersService "github.com/re-worthy/backend-go/internal/handlers/services/users"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

func SetupRoutes(mux *http.ServeMux, h *handlers.TBaseHandler) {
	// auth
	mux.HandleFunc("POST /auth/register", handlers.Adapter(authService.RegisterHandler, h))
	mux.HandleFunc("POST /auth/login", handlers.Adapter(authService.LoginHandler, h))

	// users
	mux.HandleFunc("GET /users/{user_id}", handlers.Adapter(usersService.GetOneHandler, h))

	// transactions
	mux.HandleFunc("POST /transactions", handlers.Adapter(transactionsService.CreateOneHandler, h))
	mux.HandleFunc("GET /transactions/recent", handlers.Adapter(transactionsService.GetRecentHandler, h))
	mux.HandleFunc("GET /transactions", handlers.Adapter(transactionsService.GetPaginatedHandler, h))
}
