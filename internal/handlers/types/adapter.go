package handlers

import (
	"encoding/json"
	"log"
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
				log.Printf("Error validating json body:\n\t%s", errValidateBody.Error())
				w.WriteHeader(http.StatusUnprocessableEntity)
				return
			}
		}

		response, errHandleFunc := handlerFunc(r, w, reqBodyData, generalHandler)
		if errHandleFunc != nil {
			log.Printf("Error evaling handlefunc:\n\t%s", errHandleFunc.Err.Error())
			w.WriteHeader(errHandleFunc.Status_code)
			w.Write([]byte(errHandleFunc.User_err.Error()))
			return
		}

		message, errJson := json.Marshal(response)
		if errJson != nil {
			log.Printf("Error marshaling json body:\n\t%s", errJson.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errJson.Error()))
			return
		}

		log.Printf("OK: route=%s", r.URL)
		w.WriteHeader(http.StatusOK)
		w.Write(message)
	}
}
