package services

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
)

func init() {
  HelloWorldService.HelloWhoHandler = func (r *http.Request, w http.ResponseWriter, body *dto.THelloWorldRq) (error, *dto.THelloWorld) {
    return nil, &dto.THelloWorld{Hello: body.Name}
  }
}

