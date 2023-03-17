package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func InitDates() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
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
	var artist []Artist
	err = json.Unmarshal([]byte(body), &artist)
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
}
