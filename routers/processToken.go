package routers

import (
	"errors"
	"strings"

	jwti "github.com/dgrijalva/jwt-go"
	"github.com/godoquin/twittor/database"
	"github.com/godoquin/twittor/models"
)

var Email string
var IDUser string

func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("MastersDelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) == 2 {
		tk = strings.TrimSpace(splitToken[1])
		//return claims, false, "", errors.New("formato de token invalido")
	}
	//tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwti.ParseWithClaims(tk, claims, func(token *jwti.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, _ := database.CheckUserExist(claims.Email)
		if find {
			Email = claims.Email
			IDUser = claims.ID.Hex()
			return claims, find, IDUser, nil

		}
		if !tkn.Valid {
			return claims, false, "", errors.New("token invalido")
		}

	}
	return claims, false, "", err
}
