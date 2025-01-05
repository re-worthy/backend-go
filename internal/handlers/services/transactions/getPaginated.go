package transactions

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
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
		UserID:             payload.ID,
		UseTags:            0,
		UserId2:            payload.ID,
		CommaSeparatedTags: []string{},
		UseMinCreatedAt:    0,
		MinCreatedAt:       0,
		UseMaxCreatedAt:    0,
		MaxCreatedAt:       0,
		UseDescriptionWk:   0,
		DescriptionWk:      "",
		Limit:              int64(limit),
		Offset:             int64(offset),
	})
	if getTrsErr != nil {
		return nil, &handlers.ResponseError{
			Err:         getTrsErr,
			User_err:    errors.New("Cant get transactions"),
			Status_code: http.StatusInternalServerError,
		}
	}

	resp := []dto.TTransactionWTagsRs{}
	for _, tr := range trs {
		fmt.Printf("%v", tr.Text)
		resp = append(resp, dto.TTransactionWTagsRs{
			TTransactionRs: dto.TTransactionRs{
				Description: tr.Description,
				Currency:    tr.Currency,
				ID:          tr.ID,
				OwnerID:     tr.OwnerID,
				Amount:      tr.Amount,
				IsIncome:    tr.IsIncome,
				Createdat:   tr.CreatedAt,
			},
			Tags: []string{},
			// Tags: strings.Split(tr.GroupConcat, ","),
		})
	}

	return &resp, nil
}
