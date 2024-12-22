package helloworld

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type THelloWorldHandler = handlers.THandlerFunc[interface{}, dto.THelloWorld]

var HelloWorldHandler THelloWorldHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (error, *dto.THelloWorld) {
	return nil, &dto.THelloWorld{Hello: "world"}
}
