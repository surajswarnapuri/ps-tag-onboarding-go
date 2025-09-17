package user

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/save", handler.Save).Methods("POST")
	router.HandleFunc("/find/{id}", handler.Find).Methods("GET")
}
