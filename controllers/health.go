package controllers

import (
	"net/http"

	"github.com/AnonO6/contact-mail-go/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}