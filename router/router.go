package router

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
	"trBD/internal/configuration"
	"trBD/internal/database"
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
	participantgroup := "NO"
	if experience {
		participantgroup = "YES"
	}

	// Вставка данных в базу данных
	err = insertParticipant(fio, birthdate, groupnumber, phonenumber, experience, participantgroup)

	log.Printf("Registration successful")
}

func insertParticipant(fio, birthdate, groupnumber, phonenumber string, experience bool, participantgroup string) error {
	// Начинаем транзакцию
	tx, err := dbClient.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	// Подготовка SQL-запроса
	query := "INSERT INTO Participants (FIO, BirthDate, GroupNumber, PhoneNumber, Experience, ParticipantGroup) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = tx.Exec(context.Background(), query, fio, birthdate, groupnumber, phonenumber, experience, participantgroup)
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
	tmpl, err := template.ParseFiles("templates/authPage.html")
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

func OrgPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/orgPage.html")
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
	tmpl, err := template.ParseFiles("templates/finalResults.html")
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
