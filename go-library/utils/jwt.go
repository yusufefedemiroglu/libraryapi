package utils

import (
	"go-library/models"
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

var user models.User

func GenerateKey() (string, error) {

	fmt.Println(user.ID)
	// Token'ı oluştur
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 96).Unix(),
	})

	// Token'ı gizli anahtar ile imzala
	tokenString, err := token.SignedString([]byte("hasanhuseyin"))
	if err != nil {
		return "", err // Hata durumunda geri döndür
	}

	return tokenString, nil // Başarı durumunda token'ı döndür
}
