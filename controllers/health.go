package controllers

import (
	"net/http"

	"github.com/AnonO6/contact-mail-go/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"Message": "How to use: https://github.com/AnonO6/contact-mail-go"})
}