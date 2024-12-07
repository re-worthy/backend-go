package handlers

import "net/http"

type THandlerFunc[Rq any, Rs any] func(r *http.Request, w http.ResponseWriter, reqBodySchema *Rq) (error, *Rs)

