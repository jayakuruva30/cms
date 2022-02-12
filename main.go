package main

import (
	"cms/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routers := mux.NewRouter()
	routers.HandleFunc("/api/v1/countries", routes.GetCountries).Methods("GET")
	routers.HandleFunc("/api/v1/countries", routes.CreateCountries).Methods("POST")
	http.ListenAndServe(":8000", routers)

}
