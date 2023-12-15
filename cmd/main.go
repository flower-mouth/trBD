package main

import (
	"log"
	"net/http"
	"trBD/router"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/home", router.HomePage)
	mux.HandleFunc("/authPage/", router.AuthPage)
	mux.HandleFunc("/regPage/", router.RegisterHandler)
	mux.HandleFunc("/userExp/", router.UserExpPage)
	mux.HandleFunc("/userZeroExp/", router.UserZeroExpPage)
	mux.HandleFunc("/orgPage/", router.OrgPage)
	mux.HandleFunc("/addWithoutExpPage/", router.AddWithoutExpPage)
	mux.HandleFunc("/gamesPlayedExp/", router.GamesPlayedExp)
	mux.HandleFunc("/gamesPlayedZeroExp/", router.GamesPlayedZeroExp)
	mux.HandleFunc("/pointsExp/", router.PointsExpPage)
	mux.HandleFunc("/pointsZeroExp/", router.PointsZeroExpPage)
	mux.HandleFunc("/pointsOverall/", router.TournamentResultsPage)

	log.Printf("Starting server...")
	err := http.ListenAndServe(":8181", mux)
	if err != nil {
		log.Printf("Error in lauching server: %v", err)
	}
}
