package main

import (
	"cms/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	routers := mux.NewRouter()
	routers.HandleFunc("/api/v1/countries", routes.GetCountries).Methods("GET")
	routers.HandleFunc("/api/v1/countries", routes.CreateCountries).Methods("POST")
	http.ListenAndServe(":8000", routers)

}
