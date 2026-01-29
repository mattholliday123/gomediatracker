package main

import (
	"log"
	"net/http"
)

//Change status for media
func changeStatus()

/*---Video Games---*/
//handles searching for game api and display results
func searchgameHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
}

func addGameToCollection()

/*---books---*/
//handles searching for book api and display results
func searchbookHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
}

func addBookToCollection()

/*---music---*/
//handles searching for music api and display results
func searchmusicHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
}

func addMusicToCollection()

/*---movies---*/
//handles searching for movie api and display results
func searchmovieHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
}

func addMovieToCollection()


func main() {
	http.HandleFunc("/searchgame", searchgameHandler)
	http.HandleFunc("/searchbook", searchbookHandler)
	http.HandleFunc("/searchmovie", searchmovieHandler)
	http.HandleFunc("/searchmmusic", searchmusicHandler)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
