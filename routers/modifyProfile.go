package routers

import (
	"encoding/json"
	"net/http"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = database.ModifyProfileBD(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
