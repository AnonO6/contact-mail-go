package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
func getToken(w http.ResponseWriter, r *http.Request){
	var creds Credentials;
	err:= json.NewDecoder(r.Body).Decode(&creds);
	if err!= nil{
		respondWithError(w,http.StatusBadRequest,"Invalid format")
		return
	}
	if creds.Secret != os.Getenv("SECRET") {
		respondWithError(w,http.StatusUnauthorized,"Unauthorised")
		return;
	}

	expirationTime := time.Hour * 2;
	claims := &jwt.RegisteredClaims{
		Issuer:    "mailservice",
		Subject:	"123",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);
	signedToken,err := token.SignedString(jwtKey);
	if err!= nil{
		respondWithError(w, http.StatusInternalServerError,"Failed to generate token");
		return;
	}
	response := struct{
		Token string `json:"token"`
	}{
		Token: signedToken,
	}
	respondWithJSON(w,http.StatusAccepted,response);
}
