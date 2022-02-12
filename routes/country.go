package routes

import (
	"cms/controllers"
	"cms/models"
	"encoding/json"
	"net/http"
)

func GetCountries(res http.ResponseWriter, req *http.Request) {
	var response models.Response
	result, err := models.GetCountries()
	if err != nil {
		response.Error = true
	}
	response.Data = result
	response.Status = http.StatusOK
	json.NewEncoder(res).Encode(response)
}
func CreateCountries(res http.ResponseWriter, req *http.Request) {
	result, err := controllers.CreateCountry(req)
	var response models.Response
	if err != nil {
		response.Error = true
	}
	response.Data = result
	response.Status = http.StatusOK
	json.NewEncoder(res).Encode(response)

}
