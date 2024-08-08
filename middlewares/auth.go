package middlewares

import (
	"fmt"
	"net/http"

	"github.com/AnonO6/contact-mail-go/config"
	"github.com/AnonO6/contact-mail-go/models"
	"github.com/AnonO6/contact-mail-go/utils"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signedToken := r.Header.Get("Authorization")

		token, err := jwt.ParseWithClaims(signedToken, &models.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized, token invalid")
			return
		}
		fmt.Println(token.Claims);
		claims, ok := token.Claims.(*models.CustomClaims);
		if !ok {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized, Invalid claims format")
			return
		}
		r.Header.Set("SMTPUser", claims.SMTPUser)
		r.Header.Set("SMTPPass", claims.SMTPPass)

		next.ServeHTTP(w, r)
	}
}
