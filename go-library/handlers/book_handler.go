package handlers

import (
	"context"
	"fmt"
	"go-library/database"
	"go-library/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// UserID'nin nil olup olmadığını kontrol et
	if book.UserID == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}

	// MongoDB'deki books koleksiyonuna erişim
	bookCollection := database.GetCollection("books")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Kullanıcı kitaplarını bulma
	filter := bson.M{"userid": book.UserID}
	var userBooks []models.Book
	cursor, err := bookCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch user books"})
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &userBooks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error parsing user books"})
	}

	// Kullanıcı kitaplarının sayısını kontrol et
	if len(userBooks) >= 5 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "User already has 5 books",
		})
	}

	// Yeni kitabı ekleme işlemi
	result, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Book could not be created"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
		"bookId":  result.InsertedID,
	})
}

func GetBooks(c *fiber.Ctx) error {
	// Frontend'den gelen User ID (bu artık dinamik olacak)
	userID := c.Get("x-user-id") // Burada header parametre olarak gönderildiğini varsayıyoruz.

	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User ID is required",
		})
	}

	// Authorization header'ından token'ı alıyoruz
	tokenString := c.Get("Authorization")
	if len(tokenString) <= 7 || tokenString[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	// Token'ı doğrulama
	tokenString = tokenString[7:] // "Bearer " kısmını at
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		fmt.Println("Token validated successfully.")
		// Burada secret key ile token'ı doğruluyoruz
		return []byte("hasanhuseyin"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
	objID, err := primitive.ObjectIDFromHex(userID) // burda objectidye donustusturuyoz bunu yapmadigim icin 512512521521 saat ugrastim unutma bir daha :thumbsup:
	fmt.Println(err)
	// Veritabanından User ID'ye göre kitapları alıyoruz
	cursor, err := database.BookCollection.Find(context.Background(), map[string]interface{}{"userid": objID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kitaplar alınamadı",
		})

	}

	var books []models.Book
	if err := cursor.All(context.Background(), &books); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kitapları işlerken hata oluştu",
		})
	}
	return c.JSON(books)
}
