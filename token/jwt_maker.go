package token

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{
		secretKey,
	}
}

func (m *JWTMaker) CreateToken(userID, email, username string, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(userID, email, username, duration)
	if err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate access token: %v", err)
	}

	return tokenStr, claims, nil
}

func (m *JWTMaker) VerifyToken(signedToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return []byte(m.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

func (maker *JWTMaker) GetTokenFromHeader(r *http.Request) (string, error) {
	// authHeader := r.Header.Get("Authorization")
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return "", err
	}

	if cookie.Value == "" {
		return "", fmt.Errorf("no authorization header")
	}

	// fields := strings.Fields(authHeader)
	// if len(fields) != 2 || fields[0] != "Bearer" {
	// 	return "", fmt.Errorf("invalid authorization header format")
	// }

	return cookie.Value, nil
}

// func (m *JWTMaker) RefreshToken(w http.ResponseWriter, r *http.Request) {
// 	var body models.RefreshTokenReq
// 	if err := helpers.BindValidateJSON(w, r, &body); err != nil {
// 		return
// 	}
// 	if body.RefreshToken == "" {
// 		helpers.ErrJSONResponse(w, "refresh token is required", http.StatusBadRequest)
// 		return
// 	}

// 	refreshClaims, err := m.VerifyToken(body.RefreshToken)
// 	if err != nil {
// 		helpers.ErrJSONResponse(w, fmt.Sprintf("error verifying refresh token: %v", err.Error()), http.StatusUnauthorized)
// 		return
// 	}

// 	var session models.Session
// 	if err := database.DB.QueryRowContext(
// 		r.Context(),
// 		` SELECT id, email, username, refresh_token, is_revoked, expires_at
// 		  FROM sessions
// 		  WHERE refresh_token = $1
// 		  LIMIT 1
// 		`,
// 		body.RefreshToken,
// 	).Scan(
// 		&session.ID,
// 		&session.Email,
// 		&session.Username,
// 		&session.RefreshToken,
// 		&session.IsRevoked,
// 		&session.ExpiresAt,
// 	); err != nil {
// 		if err == sql.ErrNoRows {
// 			helpers.ErrJSONResponse(w, fmt.Sprintf("error getting session: %v", "you must passs refresh token"), http.StatusInternalServerError)
// 			return
// 		}
// 		helpers.ErrJSONResponse(w, fmt.Sprintf("error getting session: %v", err.Error()), http.StatusInternalServerError)
// 		return
// 	}

// 	if session.IsRevoked {
// 		helpers.ErrJSONResponse(w, "session is revoked", http.StatusUnauthorized)
// 		return
// 	}

// 	if session.Email != refreshClaims.Email {
// 		helpers.ErrJSONResponse(w, "invalid session", http.StatusUnauthorized)
// 		return
// 	}

// 	accessToken, accessClaims, err := m.CreateToken(refreshClaims.ID, refreshClaims.Email, refreshClaims.Username, 15*time.Minute)
// 	if err != nil {
// 		helpers.ErrJSONResponse(w, fmt.Sprintf("error creating token: %v", err.Error()), http.StatusInternalServerError)
// 		return
// 	}

// 	res := models.RenewAccessTokenResponse{
// 		AccessToken:          accessToken,
// 		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
// 	}

// 	helpers.JSONResponse(w, "", res)
// }
