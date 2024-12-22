package helloworld

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/re-worthy/backend-go/internal/handlers/dto"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tHelloDBHandler = handlers.THandlerFunc[interface{}, dto.THelloDB]

var HelloDBHandler tHelloDBHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (error, *dto.THelloDB) {
	var result dto.THelloDB
	var cnt int32
	ErrNotExists := errors.New("Counter not found")
	ErrGeneralError := errors.New("Internal server error")

	_, errUpd := g.DB.Exec("UPDATE counter SET cnt=cnt+1 WHERE id=1")
	if errUpd != nil {
		return errUpd, nil
	}

	row := g.DB.QueryRow("SELECT cnt FROM counter WHERE id=1")
	errScan := row.Scan(&cnt)
	if errors.Is(errScan, sql.ErrNoRows) {
		return ErrNotExists, nil
	}
	if errScan != nil {
		log.Println(errScan.Error())
		return ErrGeneralError, nil
	}

	result.Counter = cnt

	return nil, &result
}
