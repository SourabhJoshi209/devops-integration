package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

// album represents data about a record album.
type album struct {
        ID     string  `json:"id"`
        Title  string  `json:"title"`
        Artist string  `json:"artist"`
        Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
        {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
        {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
        {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func (c *album) IsEmpty() bool {
	return c.Title == ""
}

func main() {
	fmt.Println("API - Demo")
	r := mux.NewRouter()
		r.HandleFunc("/", hello).Methods("GET")
		r.HandleFunc("/albums", getAlbums).Methods("GET")
		r.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
		r.HandleFunc("/albums", postAlbums).Methods("POST")
        
		log.Fatal(http.ListenAndServe(":4000", r))
}


func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>API - Demo</h1>"))
}
// getAlbums responds with the list of all albums as JSON.
func getAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Albums")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one album")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through albums, find matching id and return the response
	for _, album := range albums {
		if album.ID == params["id"] {
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	json.NewEncoder(w).Encode("No album found with given id")
	return
}

func postAlbums(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "applicatioan/json")

	// if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}


	var newAlbum album
	_ = json.NewDecoder(r.Body).Decode(&newAlbum)
	if newAlbum.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	newAlbum.ID = strconv.Itoa(rand.Intn(100))
	albums = append(albums, newAlbum)
	json.NewEncoder(w).Encode(albums)
	return

}














