package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AnonO6/contact-mail-go/config"
	"github.com/AnonO6/contact-mail-go/models"
	"github.com/AnonO6/contact-mail-go/utils"
	"github.com/golang-jwt/jwt/v5"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid format")
		return
	}
	if creds.Secret != os.Getenv("SECRET") {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	//Keeping a very small token alive time, to minimize risk
	expirationTime := time.Minute*15;
	fmt.Println(creds);
	claims := &models.CustomClaims{
		SMTPUser: creds.SMTPUser,
		SMTPPass: creds.SMTPPass,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "contact_mail_go",
			Subject:   "User",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.JwtKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}
	response := struct {
		Token string `json:"token"`
	}{
		Token: signedToken,
	}
	utils.RespondWithJSON(w, http.StatusAccepted, response)
}
