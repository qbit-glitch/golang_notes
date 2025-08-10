package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"school_management_api/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

func JWTMiddleware(next http.Handler) http.Handler {
	fmt.Println("---------- JWT Middleware ---------")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("++++++++++ Inside JWT Middleware +++++++++")

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte slice containing your secret, eg: []byte("my_secret_key")
			return []byte(jwtSecret), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token Expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed){
				http.Error(w, "Token Malformed", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid JWT:", token.Value)
		}
		
		fmt.Println("Parsed Token:", parsedToken)

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if ok {
			fmt.Println(claims["uid"], claims["exp"], claims["role"])
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			return
		}

		// Now use context to carry the claim information accross different middlewares, accross different functions
		ctx := context.WithValue(r.Context(), ContextKey("role"), claims["role"])
		ctx = context.WithValue(ctx, ContextKey("expiresAt"), claims["exp"])
		ctx = context.WithValue(ctx, ContextKey("username"), claims["user"])
		ctx = context.WithValue(ctx, ContextKey("userId"), claims["uid"])

		fmt.Println(ctx)
		
		next.ServeHTTP(w, r.WithContext((ctx)))
		fmt.Println("Sent Response from JWT Middleware")
	})
}
