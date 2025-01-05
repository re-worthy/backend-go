package transactions

import (
	"errors"
	"net/http"
	"strconv"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	"github.com/re-worthy/backend-go/internal/handlers/services/shared"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tGetRecentHandler = handlers.THandlerFunc[interface{}, []dto.TTransactionRs]

var GetRecentHandler tGetRecentHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (*[]dto.TTransactionRs, *handlers.ResponseError) {
	MAX_LIMIT := 3
	PARAM_NAME_LIMIT := "limit"

	token := shared.GetRequestAuth(r)
	payload, parseTokenErr := shared.GetTokenPayload(token)
	if parseTokenErr != nil {
		return nil, &handlers.ResponseError{
			Err:         parseTokenErr,
			User_err:    errors.New("Cant parse token"),
			Status_code: http.StatusUnauthorized,
		}
	}

	limit_s := r.URL.Query().Get(PARAM_NAME_LIMIT)
	limit, convertErr := strconv.Atoi(limit_s)
	if convertErr != nil {
		limit = MAX_LIMIT
	}
	limit = min(limit, MAX_LIMIT)

	trs, getTrsErr := g.Queries.GetRecentTransactionsByUserId(r.Context(), gen.GetRecentTransactionsByUserIdParams{
		OwnerID: payload.ID,
		Limit:   3,
	})
	if getTrsErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getTrsErr,
			User_err:    errors.New("Cant get transactions"),
			Status_code: http.StatusInternalServerError,
		}
	}

	resp := []dto.TTransactionRs{}
	for _, tr := range trs {
		resp = append(resp, dto.TTransactionRs{
			Description: tr.Description,
			Currency:    tr.Currency,
			ID:          tr.ID,
			OwnerID:     tr.OwnerID,
			Amount:      tr.Amount,
			IsIncome:    tr.IsIncome,
			Createdat:   tr.CreatedAt,
		})
	}

	return &resp, nil
}
