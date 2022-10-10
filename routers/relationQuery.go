package routers

import (
	"encoding/json"
	"net/http"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func RelationQuery(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID
	var resp models.ResponseRelationQuery

	status, err := database.GetRelationBD(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
