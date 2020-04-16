package router

import (
	"net/http"

	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/case", middleware.GetAllCase).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/case", middleware.CreateCase).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/officer", middleware.GetAllOfficer).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/officer", middleware.CreateOfficer).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/caseComplete", middleware.CaseComplete).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/caseComplete", middleware.CaseComplete).Methods("POST", "OPTIONS")
	router.PathPrefix("/temp-images/").Handler(http.StripPrefix("/temp-images/", http.FileServer(http.Dir("./temp-images"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./../build")))
	return router
}
