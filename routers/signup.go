package routers

import (
	"encoding/json"
	"net/http"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, find, _ := database.CheckUserExist(t.Email)
	if find {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := database.InsertRecord(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logró insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
