package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/re-worthy/backend-go/pkg/utils"
)

func Adapter[Rq any, Rs any](handlerFunc THandlerFunc[Rq, Rs]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    var reqBodyData *Rq
    reqBodyData = nil

    if r.Method != "GET" {
      var errValidateBody error
      reqBodyData, errValidateBody = utils.HttpValidateBodyJson[Rq](r, w)
      if errValidateBody != nil {
        return
      }
    }

		errHandleFunc, response := handlerFunc(r, w, reqBodyData)
		if errHandleFunc != nil {
			return
		}

		w.WriteHeader(http.StatusOK)

		message, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(message)
	}
}
