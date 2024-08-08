package models

type EmailData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Subject	string	`json:"subject"`
}
