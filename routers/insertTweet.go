package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	row := models.InsertTweet{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := database.InsertTweet(row)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertat el tweet "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
