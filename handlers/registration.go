package handlers

import (
	"html/template"
	"log"
	"net/http"
	"trBD/internal/database"
)

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
	if r.Method == http.MethodPost {
		// Получение данных из формы
		fio := r.FormValue("fio")
		birthdate := r.FormValue("birthdate")
		groupnumber := r.FormValue("groupnumber")
		phonenumber := r.FormValue("phonenumber")
		experienceValue := r.FormValue("experience")
		// Преобразуем значение из строки в булево
		experience := experienceValue == "true"

		err = database.InsertParticipant(fio, birthdate, groupnumber, phonenumber, experience)

		// Вставка данных в базу данных
		if err == nil {
			log.Printf("Registration successful")
		}
	}
}
