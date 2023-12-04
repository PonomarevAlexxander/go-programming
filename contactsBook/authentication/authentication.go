package app

import (
	"context"
	"fmt"
	"go/course/contactsbook/models"
	u "go/course/contactsbook/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/user/new", "/user/login"}
		requestPath := r.URL.Path
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			u.AuthorizationError(w, "Missing authentication token")
			return
		}

		splitted := strings.Split(tokenHeader, " ") //Bearer {token-body}
		if len(splitted) != 2 {
			u.AuthorizationError(w, "Invalid/Malformed authentication token")
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			u.AuthorizationError(w, "Malformed authentication token")
			return
		}

		if !token.Valid {
			u.AuthorizationError(w, "Token is not valid.")
			return
		}

		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
