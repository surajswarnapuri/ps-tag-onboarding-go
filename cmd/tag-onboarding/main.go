package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	userApplication "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/application/user"
	userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
	userInfra "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/infrastructure/persistence/in-memory"
	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/infrastructure/persistence/mongodb"
	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/middleware"
	userInterface "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/user"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGO_URI")
	var userRepository userEntity.Repository

	collection := os.Getenv("MONGO_COLLECTION")
	if collection == "" {
		log.Default().Println("MONGO_COLLECTION is not set, defaulting to user")
		collection = "user"
	}
	client, err := mongodb.NewMongoDBClient(ctx, uri, collection)
	if err != nil {
		log.Default().Printf("Failed to create MongoDB client: %v\n Defaulting to inmemory storage", err)
		userRepository = userInfra.NewRepository()
	} else {

		userRepository = mongodb.NewRepository(client)
	}

	userValidationService := userEntity.NewValidationService()
	userService := userApplication.NewService(userValidationService, userRepository)
	userHandler := userInterface.NewHandler(userService)

	// HTTP Server Setup
	mux := mux.NewRouter()

	// Middleware Setup - Apply before routes
	mux.Use(middleware.RequestLogger)

	userHandler.Find().AddRoute(mux)
	userHandler.Save().AddRoute(mux)

	log.Println("Starting HTTP server on port 8080")
	http.ListenAndServe(":8080", mux)

}
