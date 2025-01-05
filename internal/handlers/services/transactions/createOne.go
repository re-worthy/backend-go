package transactions

import (
	"errors"
	"net/http"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	"github.com/re-worthy/backend-go/internal/handlers/dto"
	"github.com/re-worthy/backend-go/internal/handlers/services/shared"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

type tCreateOneHandler = handlers.THandlerFunc[dto.TTransactionRq, dto.TTransactionWTagsRs]

var CreateOneHandler tCreateOneHandler = func(r *http.Request, w http.ResponseWriter, body *dto.TTransactionRq, g *handlers.TBaseHandler) (*dto.TTransactionWTagsRs, *handlers.ResponseError) {
	token := shared.GetRequestAuth(r)
	payload, parseTokenErr := shared.GetTokenPayload(token)
	if parseTokenErr != nil {
		return nil, &handlers.ResponseError{
			Err:         parseTokenErr,
			User_err:    errors.New("Cant parse token"),
			Status_code: http.StatusUnauthorized,
		}
	}

	tr, createTrErr := g.Queries.CreateTransaction(r.Context(), gen.CreateTransactionParams{
		Description: body.Description,
		Currency:    "BYN",
		OwnerID:     payload.ID,
		Amount:      body.Amount,
		IsIncome:    body.IsIncome,
	})
	if createTrErr != nil {
		return nil, &handlers.ResponseError{
			Err:         createTrErr,
			User_err:    errors.New("Cant create transaction"),
			Status_code: http.StatusInternalServerError,
		}
	}

	/*
		iTags := make([]interface{}, len(body.Tags))
		for i, s := range body.Tags {
			iTags[i] = s
		}
		iUIDs := make([]interface{}, len(body.Tags))
		for i := range body.Tags {
			iUIDs[i] = payload.ID
		}
		iTIDs := make([]interface{}, len(body.Tags))
		for i := range body.Tags {
			iTIDs[i] = tr.ID
		}
		tags, createTagsErr := g.Queries.CreateTagsBatch(r.Context(), gen.CreateTagsBatchParams{
			Texts:          iTags,
			UserIds:        iUIDs,
			TransactionIds: iTIDs,
		})
		if createTagsErr != nil {
			return nil, &handlers.ResponseError{
				Err:         createTagsErr,
				User_err:    errors.New("Cant create tags"),
				Status_code: http.StatusInternalServerError,
			}
		}
	*/

	tags := []gen.Tag{}
	for _, v := range body.Tags {
		tag, createTagErr := g.Queries.CreateTag(r.Context(), gen.CreateTagParams{
			Text:          v,
			UserID:        payload.ID,
			TransactionID: tr.ID,
		})
		if createTagErr != nil {
			return nil, &handlers.ResponseError{
				Err:         createTagErr,
				User_err:    errors.New("Cant create tag"),
				Status_code: http.StatusInternalServerError,
			}
		}

		tags = append(tags, tag)
	}

	respTags := []string{}
	for _, s := range tags {
		respTags = append(respTags, s.Text)
	}
	return &dto.TTransactionWTagsRs{
		TTransactionRs: dto.TTransactionRs{
			Description: tr.Description,
			Currency:    tr.Currency,
			ID:          tr.ID,
			OwnerID:     tr.OwnerID,
			Amount:      tr.Amount,
			IsIncome:    tr.IsIncome,
			Createdat:   tr.CreatedAt,
		},
		Tags: respTags,
	}, nil
}
