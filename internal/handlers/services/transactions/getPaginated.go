package transactions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	gen "github.com/re-worthy/backend-go/internal/db/sqlrc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	"github.com/re-worthy/backend-go/internal/handlers/services/shared"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tGetPaginatedHandler = handlers.THandlerFunc[interface{}, []dto.TTransactionWTagsRs]

var GetPaginatedHandler tGetPaginatedHandler = func(r *http.Request, w http.ResponseWriter, body *interface{}, g *handlers.TBaseHandler) (*[]dto.TTransactionWTagsRs, *handlers.ResponseError) {
	PARAM_NAME_LIMIT := "limit"
	PARAM_NAME_OFFSET := "offset"
	MAX_LIMIT := 100

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
	limit, convertLimitErr := strconv.Atoi(limit_s)
	if convertLimitErr != nil || limit > MAX_LIMIT {
		return nil, &handlers.ResponseError{
			Err:         convertLimitErr,
			User_err:    errors.New(fmt.Sprintf("Invalid limit value. Provide int <= %d", MAX_LIMIT)),
			Status_code: http.StatusUnprocessableEntity,
		}
	}

	offset_s := r.URL.Query().Get(PARAM_NAME_OFFSET)
	offset, convertOffsetErr := strconv.Atoi(offset_s)
	if convertOffsetErr != nil {
		return nil, &handlers.ResponseError{
			Err:         convertOffsetErr,
			User_err:    errors.New("Invalid offset value. Provide int"),
			Status_code: http.StatusUnprocessableEntity,
		}
	}

	toLog := map[string]any{
		"limit":  limit,
		"offset": offset,
	}

	log.Printf("%v", toLog)

	/*
		  return nil, &handlers.ResponseError{
				Err:         errors.New("wait for sqlx"),
				User_err:    errors.New("wait for sqlx"),
				Status_code: http.StatusTooEarly,
			}
	*/

	trs, getTrsErr := g.Queries.GetTransactionsByAndTags(r.Context(), gen.GetTransactionsByAndTagsParams{
		User_id:        payload.ID,
		Tags:           "",
		User_id2:       payload.ID,
		Min_created_at: 0,
		Max_created_at: 0,
		Description_wk: "",
		Limit:          limit,
		Offset:         offset,
	})
	if getTrsErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getTrsErr,
			User_err:    errors.New("Cant get transactions"),
			Status_code: http.StatusInternalServerError,
		}
	}

	resp := []dto.TTransactionWTagsRs{}
	for _, tr := range *trs {
		fmt.Printf("%v", tr.Text)
		resp = append(resp, dto.TTransactionWTagsRs{
			TTransactionRs: dto.TTransactionRs{
				Description: tr.Description,
				Currency:    tr.Currency,
				ID:          tr.Id,
				OwnerID:     tr.Owner_id,
				Amount:      tr.Amount,
				IsIncome:    tr.Is_income,
				Createdat:   tr.Created_at,
			},
			Tags: []string{},
			// Tags: strings.Split(tr.GroupConcat, ","),
		})
	}

	return &resp, nil
}
