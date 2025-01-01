package users

import (
	"net/http"
	"strconv"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tGetOneHandler = handlers.THandlerFunc[interface{}, dto.TGetUserRs]

var GetOneHandler tGetOneHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (error, *dto.TGetUserRs) {
	id_string := r.PathValue("user_id")
	id, convertErr := strconv.Atoi(id_string)
	if convertErr != nil {
		return convertErr, nil
	}

	user, getUserErr := g.Queries.GetUser(r.Context(), int64(id))
	if getUserErr != nil {
		return getUserErr, nil
	}

	return nil, &dto.TGetUserRs{
		Username: user.Username,
		Balance:  user.Balance,
		Id:       user.ID,
		Image:    user.Image,
	}
}
