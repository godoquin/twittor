package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var file_ string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(file_, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	var status bool
	user.Banner = IDUser + "." + extension
	status, err = database.ModifyProfileBD(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
