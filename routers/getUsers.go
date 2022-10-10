package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/godoquin/twittor/database"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page como entero mayor a 0", http.StatusBadRequest)
	}
	pag := int64(pagTemp)
	result, status := database.GetAllUsersBD(IDUser, pag, search, typeUser)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
