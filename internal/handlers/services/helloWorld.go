package services

import "net/http"

type THelloWorld struct {
	Hello string `json:"hello"`
}
type THelloWorldR struct {
	Name string `json:"name" validate:"required" `
}

func HelloWorldHandler(r *http.Request, w http.ResponseWriter, body THelloWorldR) (error, *THelloWorld) {
	return nil, &THelloWorld{Hello: body.Name}
}
