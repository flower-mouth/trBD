package main

import (
	"log"
	"net/http"
	"trBD/router"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router.HomePage)
	mux.HandleFunc("/authPage/", router.AuthPage)
	mux.HandleFunc("/regPage/", router.RegPage)
	mux.HandleFunc("/orgPage/", router.OrgPage)
	mux.HandleFunc("/intermediateResults/", router.IntermediateResults)
	mux.HandleFunc("/finalResults/", router.FinalResults)

	log.Printf("Starting server...")
	err := http.ListenAndServe(":8181", mux)
	if err != nil {
		log.Printf("Error in lauching server: %v", err)
	}
}
