package app

import (
	"encoding/json"
	"log"
	"net/http"
	"timezone-api/domain"
	"timezone-api/error"
	"timezone-api/service"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	timeRepository := domain.NewDefaultTimeRepository()
	timeService := service.NewDefaultTimeService(timeRepository)
	timeHandler := TimeHandler{timeService}

	router.HandleFunc("/api/time", timeHandler.handleTimeAPI)
	router.PathPrefix("/").HandlerFunc(handleUndefinedEndpoint)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func handleUndefinedEndpoint(w http.ResponseWriter, r *http.Request) {
	err := error.NewUndefinedEndpointError("endpoint not found")
	writeResponse(w, err.Code, err.AsMessage())
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
