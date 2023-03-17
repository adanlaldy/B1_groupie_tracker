package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Index struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func InitRelation() []Index {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relations")
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
	var index []Index
	err = json.Unmarshal([]byte(body), &index)
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
	return index
}
