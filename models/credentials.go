package models

type Credentials struct {
	Secret   string `json:"secret"`
	SMTPUser string `json:"smtp_user"`
	SMTPPass string `json:"smtp_pass"`
}
