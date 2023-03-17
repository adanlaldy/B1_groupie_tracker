package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func InitAllArtist() []Artist {
	var artist []Artist
	for i := 1; i <= 52; i++ {
		artist = append(artist, InitArtist(i))
	}
	return artist
}
func InitArtist(ID int) Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(ID))
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
	var artist Artist
	err = json.Unmarshal([]byte(body), &artist)
	if err != nil {
		// Gestion d'erreur
		panic(err)
	}
	return artist
}
