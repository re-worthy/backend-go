package users

import (
	"errors"
	"net/http"
	"net/url"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tCreateOneHandler = handlers.THandlerFunc[dto.TCreateUserRq, dto.TGetUserRs]

var CreateHandler tCreateOneHandler = func(r *http.Request, w http.ResponseWriter, body *dto.TCreateUserRq, g *handlers.TBaseHandler) (*dto.TGetUserRs, *handlers.ResponseError) {
	result_image := body.Image
	if result_image == "" {
		result_image = "https://api.dicebear.com/7.x/identicon/svg?seed=" + url.QueryEscape(body.Username)
	}

	// TODO add password hash
	result_password := body.Password

	user, createUserErr := g.Queries.CreateUser(r.Context(), gen.CreateUserParams{
		Username: body.Username,
		Password: result_password,
		Image:    result_image,
	})

	if createUserErr != nil {
		return nil, &handlers.ResponseError{
			Err:         createUserErr,
			User_err:    errors.New("Cant create user"),
			Status_code: http.StatusBadRequest,
		}
	}

	return &dto.TGetUserRs{
		Username: user.Username,
		Balance:  user.Balance,
		Id:       user.ID,
		Image:    user.Image,
	}, nil
}
