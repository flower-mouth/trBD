package main

import (
	"log"
	"net/http"
	"trBD/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/home", handlers.HomePage)
	mux.HandleFunc("/authPage/", handlers.AuthPage)
	mux.HandleFunc("/regPage/", handlers.RegisterHandler)
	mux.HandleFunc("/userExp/", handlers.UserExpPage)
	mux.HandleFunc("/userZeroExp/", handlers.UserZeroExpPage)
	mux.HandleFunc("/orgPage/", handlers.OrgPage)
	mux.HandleFunc("/addWithoutExpPage/", handlers.AddWithoutExpPage)
	mux.HandleFunc("/gamesPlayedExp/", handlers.GamesPlayedExp)
	mux.HandleFunc("/gamesPlayedZeroExp/", handlers.GamesPlayedZeroExp)
	mux.HandleFunc("/pointsExp/", handlers.PointsExpPage)
	mux.HandleFunc("/pointsZeroExp/", handlers.PointsZeroExpPage)
	mux.HandleFunc("/pointsOverall/", handlers.TournamentResultsPage)

	log.Printf("Starting server...")
	err := http.ListenAndServe(":8181", mux)
	if err != nil {
		log.Printf("Error in lauching server: %v", err)
	}
}
