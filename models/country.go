package models

import (
	"cms/config"
	"time"
)

type Countries struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	IsActive   int       `json:"is_active"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

func CreateCountry(country Countries) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}
	insertRecord := Countries{Name: country.Name, Created_At: time.Now()}
	db.Create(&insertRecord)
	return nil
}
func GetCountries() ([]Countries, error) {
	var countries []Countries
	db, err := config.GetDB()
	if err != nil {
		return countries, err
	}
	db.Find(&countries)
	return countries, nil
}
