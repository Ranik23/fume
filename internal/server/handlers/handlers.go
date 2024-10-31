package handlers

import (
	"encoding/json"
	sl "fume/internal/lib/logger"
	"fume/internal/repository"
	"net/http"
	"strings"
)

type RequestBody struct {
	Text string `json:"text"`
}

func SubmitHandler(repository *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": strings.Join(repository.Run(requestBody.Text), " ")}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
			return
		}

		// if err := repository.DataBase.CreateRequest(r); err != nil {
		// 	repository.Logger.Error("Ошибка при создании запроса в базе данных", sl.Err(err))
		// 	http.Error(w, "Ошибка при обработке запроса", http.StatusInternalServerError)
		// 	return
		// }
	}
}

func ConfigHandler(repository *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		config := map[string]interface{}{
			"alphabet":   repository.Config.Alphabet,
			"dictionary": repository.Config.Dictionary,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(config); err != nil {
			http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
			return
		}

		// if err := repository.DataBase.CreateRequest(r); err != nil {
		// 	repository.Logger.Error("Ошибка при создании запроса в базе данных", sl.Err(err))
		// 	http.Error(w, "Ошибка при обработке запроса", http.StatusInternalServerError)
		// 	return
		// }
	}
}

func ShowLogin(repository *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/home/anton/rosl/internal/static/login.html")

		if err := repository.DataBase.CreateRequest(r); err != nil {
			repository.Logger.Error("Ошибка при создании запроса в базе данных", sl.Err(err))
			http.Error(w, "Ошибка при обработке запроса", http.StatusInternalServerError)
			return
		}
	}
}

func HandleLogin(repository *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "password" {
			http.ServeFile(w, r, "/home/anton/rosl/internal/static/index.html")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
		}

		// if err := repository.DataBase.CreateRequest(r); err != nil {
		// 	repository.Logger.Error("Ошибка при создании запроса в базе данных", sl.Err(err))
		// 	http.Error(w, "Ошибка при обработке запроса", http.StatusInternalServerError)
		// 	return
		// }
	}
}
