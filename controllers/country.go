package controllers

import (
	"cms/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateCountry(req *http.Request) (interface{}, error) {
	var countries models.Countries
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}
	err = models.CreateCountry(countries)
	if err != nil {
		return nil, err
	}
	return "Record inserted successfully", nil
}
