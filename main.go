package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// logger is middleware that logs each request: method, path, remote addr, and duration.
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
//Change status for media
//func changeStatus()

/*---Video Games---*/
//handles searching for game api and display results; supports limit and offset for "see more"
func searchgameHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load("keys.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientid := os.Getenv("clientid")
	auth := os.Getenv("accesstoken")
	query := r.URL.Query().Get("q")
	limitNum := 10
	if l := r.URL.Query().Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n >= 1 && n <= 50 {
			limitNum = n
		}
	}
	offsetNum := 0
	if o := r.URL.Query().Get("offset"); o != "" {
		if n, err := strconv.Atoi(o); err == nil && n >= 0 {
			offsetNum = n
		}
	}
	body := "fields name, summary, release_dates.human, release_dates.date, involved_companies.developer, involved_companies.company.name; search \"" + query + "\"; limit " + strconv.Itoa(limitNum) + "; offset " + strconv.Itoa(offsetNum) + ";"
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
	id := r.URL.Query().Get("gameId")
	name := r.URL.Query().Get("name")
	status := r.URL.Query().Get("status")
	date := r.URL.Query().Get("date")
	studio := r.URL.Query().Get("dev")
	db, err := sql.Open("sqlite3", "./media")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO games(id, title, year, studio, status) VALUES(?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET status = excluded.status",id, name, date, studio, status)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted successfully")
}


//get games from db
func getGames(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("sqlite3", "./media")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

 	rows, err := db.Query("SELECT * FROM games")
    if err != nil {
				fmt.Printf("error select stmt\n")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close() 
		type Game struct {
        ID     int    `json:"id"`
        Title  string `json:"title"`
				ReleaseDate string `json:"year"`
				Developer string `json:"studio"`
        Status string `json:"status"`
    }
    
    var games []Game
    
    // Iterate through rows
    for rows.Next() {
        var game Game
        err := rows.Scan(&game.ID, &game.Title, &game.ReleaseDate, &game.Developer, &game.Status)
        if err != nil {
						fmt.Printf("error iterate rows \n")
						log.Fatal(err)
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        games = append(games, game)
    }
    
    // Set content type to JSON
    w.Header().Set("Content-Type", "application/json")
    
    // Encode and send the response
    json.NewEncoder(w).Encode(games)


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
	http.Handle("/searchgame", logger(http.HandlerFunc(searchgameHandler)))
	http.Handle("/searchbook", logger(http.HandlerFunc(searchbookHandler)))
	http.Handle("/searchmovie", logger(http.HandlerFunc(searchmovieHandler)))
	http.Handle("/searchmmusic", logger(http.HandlerFunc(searchmusicHandler)))
	http.Handle("/savegame", logger(http.HandlerFunc(addGameToCollection)))
	http.Handle("/getGames", logger(http.HandlerFunc(getGames)))
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", logger(fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
