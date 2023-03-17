package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Location struct {
	Locations []string `json:"locations"`
	//Dates string `json:"dates"`
}

func InitAllLocations() []Location {
	var location []Location
	for i := 1; i <= 52; i++ {
		location = append(location, InitLocation(i))
	}
	return location
}
func InitLocation(ID int) Location {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + strconv.Itoa(ID))
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
	defer resp.Body.Close()
	// Lecture de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
	var location Location
	err = json.Unmarshal([]byte(body), &location)
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
	return location
}
