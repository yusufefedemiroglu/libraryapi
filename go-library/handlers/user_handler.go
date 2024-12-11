package handlers

import (
	"context"
	"go-library/database"
	"go-library/models"
	"go-library/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	userCollection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User could not be created"})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func HandleLogin(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	var userFromDB models.User

	err := database.UserCollection.FindOne(context.Background(), bson.D{{Key: "username", Value: user.Username}}).Decode(&userFromDB)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error: " + err.Error(),
		})
	}

	if user.Username == userFromDB.Username && user.Password == userFromDB.Password {
		var t, _ = utils.GenerateKey()
		return c.JSON(fiber.Map{
			"token":   t,
			"id":      userFromDB.ID,
			"message": "Login successful",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid username or password",
	})
}
