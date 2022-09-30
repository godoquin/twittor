package jwt

import (
	"time"

	jwti "github.com/dgrijalva/jwt-go"
	"github.com/godoquin/twittor/models"
)

func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MastersDelDesarrollo_grupodeFacebook")
	payload := jwti.MapClaims{
		"email":       t.Email,
		"firstName":   t.FirstName,
		"lastName":    t.LastName,
		"dateOfBirth": t.DateOfBirth,
		"biography":   t.Biography,
		"location":    t.Location,
		"website":     t.Website,
		"_id":         t.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwti.NewWithClaims(jwti.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
