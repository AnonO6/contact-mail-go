package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	SMTPUser string `json:"smtp_user"`
	SMTPPass string `json:"smtp_pass"`
	jwt.RegisteredClaims
}
