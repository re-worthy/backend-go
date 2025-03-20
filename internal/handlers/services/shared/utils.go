package shared

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type TTokenPayload struct {
	ID int
}

var ERR_INVALID_TOKEN = errors.New("Invalid token")

func GetTokenPayload(token string) (*TTokenPayload, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return nil, ERR_INVALID_TOKEN
	}

	i, convertErr := strconv.ParseInt(parts[1], 10, 32)
	if convertErr != nil {
		return nil, convertErr
	}
	return &TTokenPayload{ID: int(i)}, nil
}

func GetRequestAuth(r *http.Request) string {
	if len(r.Header["Authorization"]) != 1 {
		return ""
	}
	return r.Header["Authorization"][0]
}
