package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/re-worthy/backend-go/pkg/utils"
)

func Adapter[Rq any, Rs any](handlerFunc THandlerFunc[Rq, Rs], generalHandler *TBaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBodyData *Rq
		reqBodyData = nil

		if r.Method != "GET" {
			var errValidateBody error
			reqBodyData, errValidateBody = utils.HttpValidateBodyJson[Rq](r, w)
			if errValidateBody != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}
		}

		errHandleFunc, response := handlerFunc(r, w, reqBodyData, generalHandler)
		if errHandleFunc != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errHandleFunc.Error()))
			return
		}

		message, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(message)
	}
}
