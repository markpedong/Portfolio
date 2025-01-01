package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserClaims(id, email, username string, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating token ID: %w", err)
	}

	return &UserClaims{
			Email:    email,
			UserID:   id,
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        tokenID.String(),
				Subject:   email,
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			}},
		nil
}
