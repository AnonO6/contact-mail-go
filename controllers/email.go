package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/AnonO6/contact-mail-go/models"
	"github.com/AnonO6/contact-mail-go/utils"
	"github.com/wneessen/go-mail"
)

var data models.EmailData

func SendEmail(w http.ResponseWriter, r *http.Request) {
	smtpUser := r.Header.Get("SMTPUser")
	smtpPass := r.Header.Get("SMTPPass")

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid format")
		return
	}

	verifier := emailverifier.NewVerifier().DisableDomainSuggest()
	ret, err := verifier.Verify(data.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Email address not verified")
		log.Printf("verify email address failed, error is: %v", err)
		return
	}

	if !ret.Syntax.Valid || !ret.HasMxRecords || ret.Reachable == "no" || ret.RoleAccount {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	m := mail.NewMsg()
	if err := m.From(smtpUser); err != nil {
		log.Fatalf("Failed to set FROM address %s", err)
		return
	}
	if err := m.To(data.Email); err != nil {
		log.Fatalf("Failed to set TO address %s", err)
		return
	}
	m.Subject(data.Subject);

	tpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to parse HTML template")
		log.Fatalf("Failed to parse HTML template: %s", err)
		return
	}

	if err := m.SetBodyHTMLTemplate(tpl, data); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to set HTML template mail body")
		log.Fatalf("Failed to set HTML template mail body: %s", err)
		return
	}

	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithTLSPolicy(mail.TLSMandatory), mail.WithUsername(smtpUser), mail.WithPassword(smtpPass))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create client")
		log.Fatalf("Failed to create client %s", err)
		return
	}

	if err := c.DialAndSend(m); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to send email")
		log.Fatalf("Failed to send mail: %s", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "Mail Sent")
}
