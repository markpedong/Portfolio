package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"portfolio/helpers"
	"portfolio/models"
	"portfolio/token"
	"time"
)

type authKey struct{}
type fullPath struct{}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func authMiddleware(next http.Handler, tokens *token.JWTMaker) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := verifyClaimsFromAuthHeader(r, tokens)
		if err != nil {
			helpers.ErrJSONResponse(w, fmt.Sprintf("error verifying token: %v", err.Error()), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), authKey{}, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func verifyClaimsFromAuthHeader(r *http.Request, tokenMaker *token.JWTMaker) (*token.UserClaims, error) {
	token, err := tokenMaker.GetTokenFromHeader(r)
	if err != nil {
		return nil, err
	}

	claims, err := tokenMaker.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func createStack(xs ...models.MiddleWare) models.MiddleWare {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]

			next = x(next)
		}

		return next
	}
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func logging(next http.Handler) http.Handler { // ServeHTTP is already attached on http.HandlerFunc and they have the same function signature
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		ctx := context.WithValue(r.Context(), authKey{}, "")
		next.ServeHTTP(wrapped, r.WithContext(ctx))
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(time.Now()))
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{}
		origin := r.Header.Get("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func origPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.URL.Path

		ctx := context.WithValue(r.Context(), fullPath{}, originalPath)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
