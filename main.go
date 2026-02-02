package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
	"github.com/joho/godotenv"
)


//Change status for media
//func changeStatus()

/*---Video Games---*/
//handles searching for game api and display results
func searchgameHandler(w http.ResponseWriter, r *http.Request){
	err := godotenv.Load("keys.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientid := os.Getenv("clientid")
	auth := os.Getenv("accesstoken")
	query := r.URL.Query().Get("q")
	body := "fields name, summary, release_dates, involved_companies.developer, involved_companies.company.name; search \"" + query +"\"; limit 5;"
	req, err := http.NewRequest(http.MethodPost, "https://api.igdb.com/v4/games", strings.NewReader(body))
	req.Header.Set("Client-ID", clientid)
	req.Header.Set("Authorization", "Bearer " + auth)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
	w.WriteHeader(http.StatusOK)
	w.Write(resp_body)
}

//adds game to Database
func addGameToCollection(w http.ResponseWriter, r *http.Request){
//id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	status := r.URL.Query().Get("status")
	date := r.URL.Query().Get("date")
	studio := r.URL.Query().Get("dev")
	db, err := sql.Open("sqlite3", "./media")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO games(title, year, studio, status) VALUES(?,?,?,?)", name, date, studio, status)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted successfully")

}

/*---books---*/
//handles searching for book api and display results
func searchbookHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	print(query)
}

//func addBookToCollection()

/*---music---*/
//handles searching for music api and display results
func searchmusicHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	print(query)
}

//func addMusicToCollection()

/*---movies---*/
//handles searching for movie api and display results
func searchmovieHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")
	print(query)
}

//func addMovieToCollection()

func main() {
	http.HandleFunc("/searchgame", searchgameHandler)
	http.HandleFunc("/searchbook", searchbookHandler)
	http.HandleFunc("/searchmovie", searchmovieHandler)
	http.HandleFunc("/searchmmusic", searchmusicHandler)
	http.HandleFunc("/savegame", addGameToCollection)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
