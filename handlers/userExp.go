package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"sort"
	"trBD/internal/database"
	"trBD/internal/models"
)

func UserExpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/userExp.html")
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

func GamesPlayedExp(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из таблицы ChessPlayersResults с участием FIO
	query := "SELECT p1.FIO AS Participant1FIO, p2.FIO AS Participant2FIO, c.PointsParticipant1, c.PointsParticipant2 FROM ChessPlayersResults c INNER JOIN Participants p1 ON c.Participant1ID = p1.ID INNER JOIN Participants p2 ON c.Participant2ID = p2.ID"
	rows, err := database.DBClient.Query(context.Background(), query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	defer rows.Close()

	var results []models.ChessPlayersResults

	for rows.Next() {
		var result models.ChessPlayersResults
		err := rows.Scan(&result.Participant1FIO, &result.Participant2FIO, &result.PointsParticipant1, &result.PointsParticipant2)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", 500)
			return
		}

		// Добавляем результат в массив
		results = append(results, result)
	}

	// Добавляем результаты в контекст шаблона
	data := struct {
		Results []models.ChessPlayersResults
	}{
		Results: results,
	}

	tmpl, err := template.ParseFiles("templates/gamesPlayedExp.html")
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

func PointsExpPage(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из базы данных
	pointsData, err := database.GetChessPlayersPointsFromDB(database.DBClient)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	// Сортируем данные по набранным очкам
	sort.Slice(pointsData, func(i, j int) bool {
		return pointsData[i].Points > pointsData[j].Points
	})

	// Передаем отсортированные данные в HTML-шаблон
	data := struct {
		PointsData []models.ChessPlayersPoints
	}{
		PointsData: pointsData,
	}

	tmpl, err := template.ParseFiles("templates/pointsExp.html")
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
