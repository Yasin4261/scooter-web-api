package services

import (
	"scoter-web-api/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 72 saat geçerli
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) // 'secret' gizli anahtarın
}
