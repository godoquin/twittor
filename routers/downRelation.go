package routers

import (
	"net/http"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func DownRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID
	status, err := database.DownRelationBD(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se logró borrar la relación "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
