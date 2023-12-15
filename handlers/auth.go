package handlers

import (
	"html/template"
	"log"
	"net/http"
)

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
		if username == "Exp" && password == "Exp" {
			http.Redirect(w, r, "/userExp/", http.StatusSeeOther)
			return
		}
		if username == "zeroExp" && password == "zeroExp" {
			http.Redirect(w, r, "/userZeroExp/", http.StatusSeeOther)
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
