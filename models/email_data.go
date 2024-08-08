package models

type EmailData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	ReceiverName    string `json:"receiver_name"`
	SenderName	string `json:"sender_name"`
	Subject	string	`json:"subject"`
}
