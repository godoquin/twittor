package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/jwt"
	"github.com/godoquin/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email inválido", 400)
		return
	}
	document, exist := database.TryLogin(t.Email, t.Password)
	if !exist {
		http.Error(w, "Email y/o contraseña inválidos", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Grabar cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
