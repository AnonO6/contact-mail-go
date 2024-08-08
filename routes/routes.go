package routes

import (
	"github.com/AnonO6/contact-mail-go/controllers"
	"github.com/AnonO6/contact-mail-go/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/getToken", controllers.GetToken)
	r.Post("/api/sendEmail", middlewares.Authenticate(controllers.SendEmail))
	r.Get("/api/healthz", controllers.Health)

	return r
}
