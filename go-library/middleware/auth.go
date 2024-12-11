package middleware

import (
	"fmt"
	"go-library/database"
	"go-library/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// AuthMiddleware validates JWT token in the Authorization header.
func RequireAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "no auth header")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hasanhuseyin"), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return fiber.NewError(fiber.StatusUnauthorized, "token expired")
		}

		var user models.User
		filter := bson.M{"_id": claims["id"]}
		err := database.UserCollection.FindOne(c.Context(), filter).Decode(&user)
		fmt.Println(err)

		return fiber.NewError(fiber.StatusUnauthorized, "user not found")
	}
	return c.Next()
}
