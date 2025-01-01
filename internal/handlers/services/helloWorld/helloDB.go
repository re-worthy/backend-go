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

var HelloDBHandler tHelloDBHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (*dto.THelloDB, error) {
	var result dto.THelloDB
	var cnt int
	ErrNotExists := errors.New("Counter not found")
	ErrGeneralError := errors.New("Internal server error")

	_, errUpd := g.DB.Exec("UPDATE counter SET cnt=cnt+1 WHERE id=1")
	if errUpd != nil {
		return nil, errUpd
	}

	row := g.DB.QueryRow("SELECT cnt FROM counter WHERE id=1")
	errScan := row.Scan(&cnt)
	if errors.Is(errScan, sql.ErrNoRows) {
		return nil, ErrNotExists
	}
	if errScan != nil {
		log.Println(errScan.Error())
		return nil, ErrGeneralError
	}

	result.Counter = cnt

	return &result, nil
}
