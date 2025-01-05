package auth

import (
	"errors"
	"net/http"
	"net/url"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tRegisterHandler = handlers.THandlerFunc[dto.TRegisterRq, dto.TAuthRs]

var RegisterHandler tRegisterHandler = func(r *http.Request, w http.ResponseWriter, body *dto.TRegisterRq, g *handlers.TBaseHandler) (*dto.TAuthRs, *handlers.ResponseError) {
	var (
		ERR_CREATE_TOKEN = errors.New("Cant create token")
		ERR_CREATE_USER  = errors.New("Cant create user")
		ERR_CREATE_HASH  = errors.New("Cant create password hash")
	)

	hash, hashErr := HashPassword(body.Password)
	if hashErr != nil {
		return nil, &handlers.ResponseError{
			Err:         hashErr,
			User_err:    ERR_CREATE_HASH,
			Status_code: http.StatusInternalServerError,
		}
	}

	img := body.Image
	if img == "" {
		img = "https://api.dicebear.com/7.x/identicon/svg?seed=" + url.QueryEscape(body.Username)
	}

	user, createUserErr := g.Queries.CreateUser(r.Context(), gen.CreateUserParams{
		Username: body.Username,
		Password: hash,
		Image:    img,
	})
	if createUserErr != nil {
		return nil, &handlers.ResponseError{
			Err:         createUserErr,
			User_err:    ERR_CREATE_USER,
			Status_code: http.StatusBadRequest,
		}
	}

	token, getTokenErr := GetToken(&user)
	if getTokenErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getTokenErr,
			User_err:    ERR_CREATE_TOKEN,
			Status_code: http.StatusInternalServerError,
		}
	}

	return &dto.TAuthRs{
		Token: token,
	}, nil
}
