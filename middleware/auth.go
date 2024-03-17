package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

type tokenClaims struct {
	jwt.MapClaims
	UserId   int    `json:"user_id"`
	UserRole string `json:"role"`
}

func Protected(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware here")
		header := r.Header[authorizationHeader]

		if header[0] == "" {
			log.Println("empty header")
			return
		}

		headerParts := strings.Split(header[0], " ")
		if len(headerParts) != 2 {
			log.Println("error header parts")
			return
		}

		userID, role, err := parseToken(headerParts[1])
		if err != nil {
			log.Println("some error parsing JWT", err)
			return
		}
		log.Println(headerParts)
		w.Header().Set("userID", strconv.Itoa(userID))
		w.Header().Set("role", role)
		next(w, r)
	}

}

func parseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.UserRole, nil
}
