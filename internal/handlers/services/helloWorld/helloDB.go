package helloworld

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tHelloDBHandler = handlers.THandlerFunc[interface{}, dto.THelloDB]

var HelloDBHandler tHelloDBHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (*dto.THelloDB, *handlers.ResponseError) {
	var result dto.THelloDB
	var cnt int

	_, errUpd := g.DB.Exec("UPDATE counter SET cnt=cnt+1 WHERE id=1")
	if errUpd != nil {
		// TODO add check for sql erorrs
		return nil, &handlers.ResponseError{
			Err:         errUpd,
			User_err:    errors.New("Cant update counter"),
			Status_code: http.StatusBadRequest,
		}
	}

	row := g.DB.QueryRow("SELECT cnt FROM counter WHERE id=1")
	errScan := row.Scan(&cnt)
	if errors.Is(errScan, sql.ErrNoRows) {
		return nil, &handlers.ResponseError{
			Err:         errScan,
			User_err:    errors.New("Counter does not exists"),
			Status_code: http.StatusBadRequest,
		}
	}
	if errScan != nil {
		return nil, &handlers.ResponseError{
			Err:         errScan,
			User_err:    errors.New("Cant get counter with id=1"),
			Status_code: http.StatusInternalServerError,
		}
	}

	result.Counter = cnt

	return &result, nil
}
