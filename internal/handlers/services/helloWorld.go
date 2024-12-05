package services

import (
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
)

func init() {
  HelloWorldService.HelloWorldHandler = func (r *http.Request, w http.ResponseWriter, body *interface{}) (error, *dto.THelloWorld) {
    return nil, &dto.THelloWorld{Hello: "world"}
  }
}

