package main

import (
	"users_service/controllers"
	_ "users_service/docs"
	"users_service/persistence"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           User Service API
// @version         1.0
// @description     Manages users and organizations data
func main() {

	persister := persistence.NewPostgrePersister()
	defer persister.Deinit()
	usersController := controllers.NewUsersController(persister)
	orgController := controllers.NewOrgController(persister)
	healthController := controllers.HealthController{}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Health
	router.Get("/health", healthController.ServeHealth)
	// Users
	router.Post("/users/create", usersController.ServeCreateUser)
	router.Post("/users/check", usersController.ServeCheckUser)
	router.Get("/users", usersController.ServeAllUsernames)
	// Orgs
	router.Post("/orgs/create", orgController.ServeCreateOrg)
	router.Get("/orgs", orgController.ServeAllOrgnames)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5101/swagger/doc.json"),
	))

	http.ListenAndServe(":5101", router)
}
