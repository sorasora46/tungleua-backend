package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
	"github.com/sorasora46/Tungleua-backend/config"
	"google.golang.org/api/option"
)

func main() {
	config.LoadConfig()

	utils.ConnectDatabase()
	config := &firebase.Config{
		ProjectID: "tungluea",
	}

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	firebaseApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}
	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Firebase Auth client: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Post("/create", handlers.Adduser)
	app.Get("/", func(c *fiber.Ctx) error {
		user := new(models.User)
		utils.DB.First(&user)

		fmt.Println(user)
		return c.JSON(user)
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		authorizationHeader := c.Get("Authorization")
		token := parseToken(authorizationHeader)
		// Perform token validation or further processing
		result, err := firebaseAuth.VerifyIDToken(context.Background(), token)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid token"})
		}
		fmt.Println("Received token: ", result.Claims["email"].(string))
		return c.SendString("Token received")
	})

	app.Post("/register", func(c *fiber.Ctx) error {

		user := new(models.User)
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}
		// Create a new user
		existingUser := user
		result := utils.DB.Where("email = ?", user.Email).First(&existingUser)
		if result.Error == nil {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "User already exists",
			})
		}

		result1 := utils.DB.Create(&user)
		if result1.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		// Create a new user in Firebase Authentication
		params := (&auth.UserToCreate{}).
			Email(user.Email).
			Password(user.Password)

		user1, err := firebaseAuth.CreateUser(context.Background(), params)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user to firebase",
			})
		}

		// Generate Firebase token
		token, err := generateFirebaseToken(firebaseAuth, "user.UID")
		if err != nil {
			log.Fatalf("Failed to generate Firebase token: %v", err)
		}

		log.Printf("Got custom token: %v\n", token)

		// Return the generated token
		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"user":    user,
			"token":   token,
			"user1":   user1,
		})
	})

	// Start the server
	app.Listen(":3000")
}

func generateFirebaseToken(client *auth.Client, uid string) (string, error) {
	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		return "", err
	}
	return token, nil
}
func parseToken(header string) string {
	// Extract the token from the Authorization header
	const bearerPrefix = "Bearer "
	if len(header) > len(bearerPrefix) && header[:len(bearerPrefix)] == bearerPrefix {
		return header[len(bearerPrefix):]
	}
	return ""
}
