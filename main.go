package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnonO6/contact-mail-go/config"
	"github.com/AnonO6/contact-mail-go/routes"
)

func main() {
	config.LoadConfig()

	r := routes.SetupRoutes()

	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
