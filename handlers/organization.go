package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"trBD/internal/database"
	"trBD/internal/models"
)

func OrgPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodPost {
		// Получение данных из формы
		firstFIO := r.FormValue("participant1")
		secondFIO := r.FormValue("participant2")

		log.Printf("firstFIO = %v", firstFIO)
		log.Printf("secondFIO = %v", secondFIO)

		firstId, err := database.GetParticipantIDByFIO(database.DBClient, firstFIO)
		if err != nil {
			log.Print(err)
		}

		secondId, err := database.GetParticipantIDByFIO(database.DBClient, secondFIO)
		if err != nil {
			log.Print(err)
		}

		log.Printf("firstId = %v", firstId)
		log.Printf("secondId = %v", secondId)

		pointsFirst, _ := strconv.ParseFloat(r.FormValue("pointsParticipant1"), 32)
		pointsSecond, _ := strconv.ParseFloat(r.FormValue("pointsParticipant2"), 32)

		log.Printf("pointsFirst = %v", pointsFirst)
		log.Printf("pointsSecond = %v", pointsSecond)

		// Вставка данных в базу данных
		err = database.InsertChessPlayerResult(firstId, secondId, pointsFirst, pointsSecond)

		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Chess game results insertion successful")
		}

		fmt.Println("----------------------------------------")
	}
	// Получаем список участников из базы данных
	participants, err := database.GetExpParticipantsFromDB(database.DBClient)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	// Добавляем участников в контекст шаблона
	data := struct {
		Participants []models.Participants
	}{
		Participants: participants,
	}

	tmpl, err := template.ParseFiles("templates/orgPage.html")
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

func AddWithoutExpPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodPost {
		// Получение данных из формы
		firstFIO := r.FormValue("participant1")
		secondFIO := r.FormValue("participant2")

		log.Printf("firstFIO = %v", firstFIO)
		log.Printf("secondFIO = %v", secondFIO)

		firstId, err := database.GetParticipantIDByFIO(database.DBClient, firstFIO)
		if err != nil {
			log.Print(err)
		}

		secondId, err := database.GetParticipantIDByFIO(database.DBClient, secondFIO)
		if err != nil {
			log.Print(err)
		}

		log.Printf("firstId = %v", firstId)
		log.Printf("secondId = %v", secondId)

		pointsFirst, _ := strconv.ParseFloat(r.FormValue("pointsParticipant1"), 32)
		pointsSecond, _ := strconv.ParseFloat(r.FormValue("pointsParticipant2"), 32)

		log.Printf("pointsFirst = %v", pointsFirst)
		log.Printf("pointsSecond = %v", pointsSecond)

		// Вставка данных в базу данных
		err = database.InsertNonChessPlayerResult(firstId, secondId, pointsFirst, pointsSecond)

		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Chess game results insertion successful")
		}

		fmt.Println("----------------------------------------")
	}
	// Получаем список участников из базы данных
	participants, err := database.GetZeroExpParticipantsFromDB(database.DBClient)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
	// Добавляем участников в контекст шаблона
	data := struct {
		Participants []models.Participants
	}{
		Participants: participants,
	}

	tmpl, err := template.ParseFiles("templates/addWithoutExpPage.html")
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
