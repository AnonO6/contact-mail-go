package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wneessen/go-mail"
)

func authenticate(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		signedToken := r.Header.Get("Authorization");

		token, err := jwt.ParseWithClaims(signedToken, &jwt.RegisteredClaims{},func(t *jwt.Token) (interface{}, error) {return jwtKey,nil});
		if err!=nil || !token.Valid {
			respondWithError(w,http.StatusUnauthorized,"Unauthorized");
			return;
		}
		next.ServeHTTP(w,r);
	}
}
//   {
// 	"email":"xyzgamer@mail",
// 	"message":"Yo!",
// 	"name":"Avi"
//   }
var data EmailData;
func validateEmail(next http.HandlerFunc ) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&data);
		if err!= nil {
			respondWithError(w,http.StatusBadRequest, "Invalid format");
			return;
		}
	// Verifying email
	verifier := emailverifier.NewVerifier().DisableDomainSuggest();
	ret, err := verifier.Verify(data.Email);
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Email address not verified")
			log.Printf("verify email address failed, error is: %v", err)
			return
		}
	   // Custom checks based on the email verification result
	   if !ret.Syntax.Valid {
        respondWithError(w, http.StatusBadRequest, "Invalid email syntax")
        return
    	}
		if !ret.HasMxRecords {
			respondWithError(w, http.StatusBadRequest, "Email domain does not have MX records")
			return
		}

		if ret.Reachable == "no" {
			respondWithError(w, http.StatusBadRequest, "Email address is not reachable")
			return
		}

		if ret.RoleAccount {
			respondWithError(w, http.StatusBadRequest, "Role-based email addresses are not allowed")
			return;
		}

		// If all checks pass, proceed to send the email
		next.ServeHTTP(w,r);
	}
}

func sendEmail(w http.ResponseWriter, r *http.Request){

	//Setting up the message body
	var email= os.Getenv("EMAIL");
	var password= os.Getenv("PASSWORD")
	m := mail.NewMsg();
	if err := m.From(email); err!=nil{
		log.Fatalf("Failed to set FROM address %s", err);
		return;
	}
	if err := m.To(data.Email); err!=nil{
		log.Fatalf("Failed to set TO address %s", err);
		return;
	}
	m.Subject("Thankyou for contacting!");

	// Parse the HTML template
	tpl, err := template.ParseFiles("email_template.html")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to parse HTML template")
		log.Fatalf("Failed to parse HTML template: %s", err)
		return;
	}
	// m.Subject("Thankyou for contacting me");
	// reply := "Hi "+ data.Name+ ",\n I read your message and will reach you back ASAP! :-D";
	// m.SetBodyString(mail.TypeTextPlain,reply);

	// Set the email body to the generated HTML content using the template and data
	if err := m.SetBodyHTMLTemplate(tpl, data); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to set HTML template mail body")
		log.Fatalf("Failed to set HTML template mail body: %s", err)
		return
	}
	
	//Setting up mail client

	c,err := mail.NewClient("smtp.gmail.com",mail.WithPort(587),mail.WithSMTPAuth(mail.SMTPAuthPlain),mail.WithTLSPortPolicy(mail.TLSMandatory),mail.WithUsername(email),mail.WithPassword(password));
	if err!= nil{
		respondWithError(w, http.StatusInternalServerError, "Failed to create client")
		log.Fatalf("Failed to create client %s",err);
		return;
	}
	//Sending the email
	if err:= c.DialAndSend(m); err!=nil{
		respondWithError(w, http.StatusInternalServerError, "Failed to send email")
		log.Fatalf("Failed to send mail: %s", err);
		return;
	}

	respondWithJSON(w,http.StatusCreated,"Mail Sent")
}