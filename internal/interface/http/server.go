package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	// Domain layer
	userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"

	// Application layer
	userApp "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/application/user"

	// Infrastructure layer
	inmemory "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/infrastructure/persistence/in-memory"

	// Interface layer
	userInterface "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/user"
)

// Server represents the HTTP server
type Server struct {
	router *mux.Router
	server *http.Server
}

// NewServer creates a new HTTP server with all dependencies wired
func NewServer() *Server {
	// Create infrastructure dependencies
	userRepo := inmemory.NewRepository()

	// Create domain services
	userValidationService := userDomain.NewValidationService()

	// Create application services
	userService := userApp.NewService(userValidationService, userRepo)

	// Create interface handlers
	userHandler := userInterface.NewHandler(userService)

	// Create router
	router := mux.NewRouter()

	// Add global middleware
	router.Use(loggingMiddleware)
	router.Use(corsMiddleware)
	router.Use(recoveryMiddleware)

	// Register user routes directly
	userInterface.RegisterRoutes(router, userHandler)

	// Add health check route
	router.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// Add a catch-all route for 404s
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	return &Server{
		router: router,
		server: &http.Server{
			Addr:         getServerAddress(),
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	log.Printf("Starting server on %s", s.server.Addr)

	// Start server in a goroutine
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
	return nil
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// Handler functions

// healthCheckHandler handles health check requests
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
		"service":   "user-service",
		"version":   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health check response: %v", err)
	}
}

// notFoundHandler handles 404 requests
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"error":   "Not Found",
		"message": "The requested resource was not found",
		"path":    r.URL.Path,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding 404 response: %v", err)
	}
}

// Helper functions

// getServerAddress returns the server address from environment or default
func getServerAddress() string {
	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		return addr
	}
	return ":8080"
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
