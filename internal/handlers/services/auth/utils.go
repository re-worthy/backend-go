package auth

import (
	"fmt"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
)

func HashPassword(password string) (string, error) {
	return password, nil
}

func ValidatePassword(input string, hash string) (bool, error) {
	isOk := input == hash

	if !isOk {
		return false, nil
	}
	return isOk, nil
}

func GetToken(user *gen.User) (string, error) {
	return fmt.Sprintf("token_for_userid:%d", user.ID), nil
}
