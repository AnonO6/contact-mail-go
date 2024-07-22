package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)
var jwtKey = []byte(os.Getenv("JWT_KEY"))
type Credentials struct {
	Secret string `json:"secret"`
}

type EmailData struct{
	Email string `json:"email"`
	Message string `json:"message"`
	Name string `json:"name"`
}
// Helper function to respond with error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
func main(){
	err := godotenv.Load();
	if err!= nil {
		log.Fatal("Error loading .env file");
		return
	}

	r := chi.NewRouter();
	r.Use(middleware.Logger);
	r.Use(middleware.Recoverer)

	r.Post("/api/getToken", getToken);
	r.Post("/api/sendEmail",authenticate(validateEmail(sendEmail)));
	fmt.Println("Server listening on port :8080");
	log.Fatal(http.ListenAndServeTLS(":8080","server.crt","server.key",r));
}