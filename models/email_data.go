package models

type EmailData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Subject	string	`json:"subject"`
}

// { 
// 	"email":"xyzgamer@gmail.com",
// 	"Message":"Just testing this",
// 	"name":"Aviral",
// 	"subject":"Hey! This is a testing message"
// }