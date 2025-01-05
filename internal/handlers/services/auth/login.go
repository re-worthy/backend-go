package auth

import (
	"errors"
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tLoginHandler = handlers.THandlerFunc[dto.TLoginRq, dto.TAuthRs]

var LoginHandler tLoginHandler = func(r *http.Request, w http.ResponseWriter, body *dto.TLoginRq, g *handlers.TBaseHandler) (*dto.TAuthRs, *handlers.ResponseError) {
	var (
		ERR_GET_USER            = errors.New("Cant find user")
		ERR_VALIDATE_PSWD       = errors.New("Cant validate password")
		ERR_CREATE_TOKEN        = errors.New("Cant create token")
		ERR_USER_INVALID_PSWD   = errors.New("Invalid username or password")
		ERR_SERVER_INVALID_PSWD = errors.New("Password validation returned false")
	)

	user, getUserErr := g.Queries.GetUserByUsername(r.Context(), body.Username)
	if getUserErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getUserErr,
			User_err:    ERR_GET_USER,
			Status_code: http.StatusBadRequest,
		}
	}

	isOk, validateErr := ValidatePassword(body.Password, user.Password)
	if validateErr != nil {
		return nil, &handlers.ResponseError{
			Err:         validateErr,
			User_err:    ERR_VALIDATE_PSWD,
			Status_code: http.StatusInternalServerError,
		}
	}
	if !isOk {
		return nil, &handlers.ResponseError{
			Err:         ERR_SERVER_INVALID_PSWD,
			User_err:    ERR_USER_INVALID_PSWD,
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
