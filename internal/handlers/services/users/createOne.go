package users

import (
	"log"
	"net/http"
	"net/url"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tCreateOneHandler = handlers.THandlerFunc[dto.TCreateUserRq, dto.TGetUserRs]

var CreateHandler tCreateOneHandler = func(r *http.Request, w http.ResponseWriter, body *dto.TCreateUserRq, g *handlers.TBaseHandler) (*dto.TGetUserRs, error) {
	result_image := body.Image
	log.Printf("result_image: %s", result_image)
	if result_image == "" {
		result_image = "https://api.dicebear.com/7.x/identicon/svg?seed=" + url.QueryEscape(body.Username)
	}
	log.Printf("result_image: %s", result_image)

	// TODO add password hash
	result_password := body.Password

	user, createUserErr := g.Queries.CreateUser(r.Context(), gen.CreateUserParams{
		Username: body.Username,
		Password: result_password,
		Image:    result_image,
	})

	if createUserErr != nil {
		return nil, createUserErr
	}

	return &dto.TGetUserRs{
		Username: user.Username,
		Balance:  user.Balance,
		Id:       user.ID,
		Image:    user.Image,
	}, nil
}
