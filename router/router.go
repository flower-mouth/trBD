package router

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"trBD/internal/configuration"
	"trBD/internal/database"
	"trBD/internal/models"
)

var dbClient database.Client

func init() {
	// Создаем клиент для работы с базой данных
	sc := configuration.StConfig{
		Username: "postgres",
		Password: "password",
		Host:     "localhost",
		Port:     "5432",
		Database: "trbd",
	}

	// Инициализируем клиент базы данных
	dbClient = database.NewClient(context.Background(), 3, sc)

	// Проверяем подключение к базе данных
	err := database.AttemptDatabaseConnection(func() error {
		conn, err := dbClient.Begin(context.Background())
		if err != nil {
			return err
		}
		defer conn.Rollback(context.Background())

		// Проверяем, что таблица Participants существует
		_, err = conn.Exec(context.Background(), "SELECT 1 FROM Participants LIMIT 1")
		if err != nil {
			return err
		}

		return nil
	}, 3, 5*time.Second)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/regPage.html")
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

	// Получение данных из формы
	fio := r.FormValue("fio")
	birthdate := r.FormValue("birthdate")
	groupnumber := r.FormValue("groupnumber")
	phonenumber := r.FormValue("phonenumber")
	experienceValue := r.FormValue("experience")
	// Преобразуем значение из строки в булево
	experience := experienceValue == "true"

	// Вставка данных в базу данных
	err = insertParticipant(fio, birthdate, groupnumber, phonenumber, experience)

	log.Printf("Registration successful")
}

func insertParticipant(fio, birthdate, groupnumber, phonenumber string, experience bool) error {
	// Начинаем транзакцию
	tx, err := dbClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "INSERT INTO Participants (FIO, BirthDate, GroupNumber, PhoneNumber, Experience) VALUES ($1, $2, $3, $4, $5)"
	_, err = tx.Exec(context.Background(), query, fio, birthdate, groupnumber, phonenumber, experience)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

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

func AuthPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получение данных из формы
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Проверка учетных данных
		if username == "admin" && password == "admin" {
			http.Redirect(w, r, "/orgPage/", http.StatusSeeOther)
			return
		}
		if username == "user" && password == "user" {
			http.Redirect(w, r, "/userPage/", http.StatusSeeOther)
			return
		}
	}

	// В противном случае, отобразить страницу входа
	tmpl, err := template.ParseFiles("templates/authPage.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func OrgPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method == http.MethodPost {
		// Получение данных из формы
		firstFIO := r.FormValue("participant1")
		secondFIO := r.FormValue("participant2")

		log.Printf("firstFIO = %v", firstFIO)
		log.Printf("secondFIO = %v", secondFIO)

		firstId, err := getParticipantIDByFIO(dbClient, firstFIO)
		if err != nil {
			log.Print(err)
		}

		secondId, err := getParticipantIDByFIO(dbClient, secondFIO)
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
		err = insertChessPlayerResult(firstId, secondId, pointsFirst, pointsSecond)

		if err != nil {
			log.Print(err)
		} else {
			log.Printf("Chess game results insertion successful")
		}
	}
	// Получаем список участников из базы данных
	participants, err := getExpParticipantsFromDB(dbClient)
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

	fmt.Println("\n")

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

func insertChessPlayerResult(participant1id, participant2id int, pointsparticipant1, pointsparticipant2 float64) error {
	// Начинаем транзакцию
	tx, err := dbClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "INSERT INTO chessplayersresults (participant1id, participant2id, pointsparticipant1, pointsparticipant2) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(context.Background(), query, participant1id, participant2id, pointsparticipant1, pointsparticipant2)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

// Получение участников из базы данных
func getExpParticipantsFromDB(dbClient database.Client) ([]models.Participants, error) {
	rows, err := dbClient.Query(context.Background(), "SELECT fio FROM Participants WHERE experience = true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.Participants

	for rows.Next() {
		var p models.Participants
		err := rows.Scan(&p.FIO)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

// Получение участников из базы данных
func getZeroExpParticipantsFromDB(dbClient database.Client) ([]models.Participants, error) {
	rows, err := dbClient.Query(context.Background(), "SELECT fio FROM Participants WHERE experience = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.Participants

	for rows.Next() {
		var p models.Participants
		err := rows.Scan(&p.FIO)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

// Функция получения ID участника по его ФИО
func getParticipantIDByFIO(dbClient database.Client, fio string) (int, error) {
	var participantID int

	err := dbClient.QueryRow(context.Background(), "SELECT ID FROM Participants WHERE FIO = $1", fio).Scan(&participantID)
	if err != nil {
		return 0, err
	}

	return participantID, nil
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/userPage.html")
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

func IntermediateResults(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/intermediateResults.html")
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

func FinalResults(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/points.html")
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
