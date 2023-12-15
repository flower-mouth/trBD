package handlers

import (
	"html/template"
	"log"
	"net/http"
	"trBD/internal/database"
	"trBD/internal/models"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/homePage.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}

func TournamentResultsPage(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из представления TournamentResults
	results, err := database.GetTournamentResultsFromDB(database.DBClient)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	// Добавляем результаты в контекст шаблона
	data := struct {
		Participants []models.TournamentResults
	}{
		Participants: results,
	}

	tmpl, err := template.ParseFiles("templates/pointsOverall.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}
