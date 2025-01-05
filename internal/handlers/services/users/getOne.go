package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tGetOneHandler = handlers.THandlerFunc[interface{}, dto.TGetUserRs]

var GetOneHandler tGetOneHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (*dto.TGetUserRs, *handlers.ResponseError) {
	id_string := r.PathValue("user_id")
	id, convertErr := strconv.Atoi(id_string)
	if convertErr != nil {
		return nil, &handlers.ResponseError{
			Err:         convertErr,
			User_err:    errors.New("Invalid user_id path param. Pass integer"),
			Status_code: http.StatusBadRequest,
		}
	}

	user, getUserErr := g.Queries.GetUserById(r.Context(), int64(id))
	if getUserErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getUserErr,
			User_err:    errors.New("User does not exist"),
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
