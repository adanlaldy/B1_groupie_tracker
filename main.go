package main

import (
	"example.com/m/models"
	"fmt"
	. "net/http"
	"strconv"
	"text/template"
)

var data Data

type Data struct {
	artist               models.Artist
	allArtist            []models.Artist
	filteredArtistsIndex []int
	filteredArtists      []models.Artist
	index                int
}

func getAllArtist() []models.Artist {
	if data.allArtist == nil {
		data.allArtist = models.InitAllArtist()
	}
	return data.allArtist
}

func setIndexFiltered(Index int) {
	data.filteredArtistsIndex = append(data.filteredArtistsIndex, Index)
}

func getIndexFiltered() []int {
	return data.filteredArtistsIndex
}

func resetIndexFiltered() {
	data.filteredArtistsIndex = data.filteredArtistsIndex[:0]
}

func initFilteredArtists() []models.Artist {
	if len(getIndexFiltered()) != 0 {
		for i := 0; i < len(getIndexFiltered()); i++ {
			data.filteredArtists = append(data.filteredArtists, getAllArtist()[getIndexFiltered()[i]])
		}
	}
	return data.filteredArtists
}

func resetFilteredArtists() {
	data.filteredArtists = data.filteredArtists[:0]
}

func setIndexArtist(mIndex int) {
	data.index = mIndex
}

func getIndexArtist() int {
	return data.index
}

func initAllArtistPages() {
	for i := 1; i <= 52; i++ {
		HandleFunc("/artist/"+strconv.Itoa(i), artistPage)
	}
}

func getNumberMembers(index int) int {
	counter := 0
	for i := 0; i < len(getAllArtist()[index].Members); i++ {
		if getAllArtist()[index].Members[i] != "" {
			counter++
		}
	}
	return counter
}

func homePage(w ResponseWriter, r *Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
	resetFilteredArtists()
	resetIndexFiltered()
	if r.FormValue("title") != "" {
		id, _ := strconv.Atoi(r.FormValue("title"))
		data.artist = models.InitArtist(id)
		Redirect(w, r, "/artist/"+strconv.Itoa(data.artist.ID), 303)
	}
	for i := 0; i < 52; i++ {
		if r.FormValue("searchbar") == getAllArtist()[i].Name {
			setIndexArtist(i)
			Redirect(w, r, "/search", 303)
		} else if r.FormValue("searchbar") == strconv.Itoa(getAllArtist()[i].CreationDate) {
			setIndexArtist(i)
			Redirect(w, r, "/search", 303)
		} else if r.FormValue("searchbar") == getAllArtist()[i].FirstAlbum {
			setIndexArtist(i)
			Redirect(w, r, "/search", 303)
		} else if r.FormValue("searchbar") == getAllArtist()[i].Locations {
			setIndexArtist(i)
			Redirect(w, r, "/search", 303)
		} else if r.FormValue("members") == strconv.Itoa(getNumberMembers(i)) {
			setIndexFiltered(i)
		}
		for j := 0; j < len(getAllArtist()[i].Members); j++ {
			if r.FormValue("searchbar") == getAllArtist()[i].Members[j] {
				setIndexArtist(i)
				Redirect(w, r, "/search", 303)
			}
		}
	}
	initFilteredArtists()
	if len(data.filteredArtistsIndex) != 0 {
		Redirect(w, r, "/searchbyfilter", 303)
	}
	t.Execute(w, data.allArtist)
}

func artistPage(w ResponseWriter, r *Request) {
	t := template.Must(template.ParseFiles("./templates/artist.html"))
	t.Execute(w, data.artist)
}

func pageNotFound(w ResponseWriter, r *Request) {
	w.WriteHeader(StatusNotFound)
	t := template.Must(template.ParseFiles("./templates/404.html"))
	t.Execute(w, nil)
}

func searchPage(w ResponseWriter, r *Request) {
	t := template.Must(template.ParseFiles("./templates/search.html"))
	if r.FormValue("title") != "" {
		id, _ := strconv.Atoi(r.FormValue("title"))
		data.artist = models.InitArtist(id)
		Redirect(w, r, "/artist/"+strconv.Itoa(data.artist.ID), 303)
	}
	t.Execute(w, data.allArtist[getIndexArtist()])
}

func filterPage(w ResponseWriter, r *Request) {
	t := template.Must(template.ParseFiles("./templates/searchbyfilter.html"))
	if r.FormValue("title") != "" {
		id, _ := strconv.Atoi(r.FormValue("title"))
		data.artist = models.InitArtist(id)
		Redirect(w, r, "/artist/"+strconv.Itoa(data.artist.ID), 303)
	}
	t.Execute(w, data.filteredArtists)
}

func contactPage(w ResponseWriter, r *Request) {
	t := template.Must(template.ParseFiles("./templates/contact.html"))
	t.Execute(w, r)
}

func main() {
	fs := FileServer(Dir("./templates"))
	Handle("/static/", StripPrefix("/static", fs))
	fmt.Println("http://localhost:8080/home")
	HandleFunc("/home", homePage)
	HandleFunc("/artist", artistPage)
	HandleFunc("/search", searchPage)
	HandleFunc("/searchbyfilter", filterPage)
	HandleFunc("/contact", contactPage)
	initAllArtistPages()
	HandleFunc("/", func(w ResponseWriter, r *Request) {
		pageNotFound(w, r)
	})
	ListenAndServe(":8080", nil)
}
