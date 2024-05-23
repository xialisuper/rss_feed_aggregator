package main

import (
	"context"
	"fmt"
	"net/http"
)

type contextKey string

const userKey = contextKey("user_key")

func (cfg *apiConfig) middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get api key from header
		key, err := getApiKeyFromHeader(r)

		if err != nil {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), key)
		if err != nil {
			http.Error(w, "user with API key not found", http.StatusUnauthorized)
			return
		}

		// set user_id in context
		ctx := r.Context()

		ctx = context.WithValue(ctx, userKey, user)
		r = r.WithContext(ctx)

		fmt.Println("Auth middleware called, user_id is: ", user.ID)

		next.ServeHTTP(w, r)
	})

}
