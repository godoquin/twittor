package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/godoquin/twittor/database"
)

func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro page", http.StatusBadRequest)
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page como entero mayor a 0", http.StatusBadRequest)
	}
	response, status := database.ReatTweetFollowers(IDUser, page)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
