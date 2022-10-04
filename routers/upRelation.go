package routers

import (
	"net/http"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func UpRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID
	status, err := database.InsertRelationBD(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar la relación ", http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la relación ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
