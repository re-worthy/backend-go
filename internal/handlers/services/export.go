package services

import (
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	"github.com/re-worthy/backend-go/internal/handlers/types"
)

type TX struct {}

type THelloService struct {
  HelloWhoHandler handlers.THandlerFunc[dto.THelloWorldRq, dto.THelloWorld]
  HelloWorldHandler handlers.THandlerFunc[interface{}, dto.THelloWorld]
}

var HelloWorldService THelloService


