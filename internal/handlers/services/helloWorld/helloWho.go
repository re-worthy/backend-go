package helloworld

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tHelloWhoHandler = handlers.THandlerFunc[dto.THelloWorldRq, dto.THelloWorld]

var HelloWhoHandler tHelloWhoHandler = func(r *http.Request, w http.ResponseWriter, body *dto.THelloWorldRq, g *handlers.TBaseHandler) (*dto.THelloWorld, *handlers.ResponseError) {
	return &dto.THelloWorld{Hello: body.Name}, nil
}
