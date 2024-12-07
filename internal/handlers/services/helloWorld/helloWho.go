package helloworld

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	types "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tHelloWhoHandler = types.THandlerFunc[dto.THelloWorldRq, dto.THelloWorld]

var HelloWhoHandler tHelloWhoHandler = func(r *http.Request, w http.ResponseWriter, body *dto.THelloWorldRq) (error, *dto.THelloWorld) {
	return nil, &dto.THelloWorld{Hello: body.Name}
}
