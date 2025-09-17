package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	userApplication "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/application/user"
	userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
	userInfra "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/infrastructure/persistence/in-memory"
	userInterface "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/user"
)

func main() {
	userRepository := userInfra.NewRepository()
	userValidationService := userEntity.NewValidationService()
	userService := userApplication.NewService(userValidationService, userRepository)
	userHandler := userInterface.NewHandler(userService)

	// HTTP Server Setup
	mux := mux.NewRouter()
	userHandler.Find().AddRoute(mux)
	userHandler.Save().AddRoute(mux)

	log.Println("Starting HTTP server on port 8080")
	http.ListenAndServe(":8080", mux)

}
