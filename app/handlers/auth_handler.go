package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	// pass := new(models.Password)
	req := new(models.LoginRequest)

	// Get request body
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	// Find email in database
	user_result := utils.DB.Find(user, "email = ?", req.Email)
	if user_result.Error != nil {
		return user_result.Error
	}

	// Find hashed password
	// pass_result := utils.DB.Find(pass, "user_id = ?", user.ID)
	// if pass_result.Error != nil {
	// 	return pass_result.Error
	// }

	// Compare password
	// if err := bcrypt.CompareHashAndPassword([]byte(pass.HashedPassword), []byte(req.Password)); err != nil {
	// 	return err
	// }

	// Create Claims (Payload)
	claims := createClaims(user)
	// Generate access token
	// access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate access token string using secret
	// acc_secret := config.GetAccessTokenSecret()
	// access_token_str, acc_sign_err := access_token.SignedString([]byte(acc_secret))
	// if acc_sign_err != nil {
	// 	return acc_sign_err
	// }

	// Generate refresh token
	// refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"ExpiresAt": jwt.NewNumericDate(time.Now().Add(3 * 24 * time.Hour)),
	// 	"id":        user.ID,
	// })
	// Generate refresh toke nstring using secret
	// rfh_secret := config.GetRefreshTokenSecret()
	// refresh_token_str, rfh_sign_err := refresh_token.SignedString([]byte(rfh_secret))
	// if rfh_sign_err != nil {
	// 	return rfh_sign_err
	// }

	opt := option.WithCredentialsFile("firebase-config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomTokenWithClaims(context.Background(), user.ID, claims)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	// return c.JSON(map[string]string{
	// 	"access_token":  access_token_str,
	// 	"refresh_token": refresh_token_str,
	// })
	return c.JSON(map[string]string{
		"token": token,
	})
}

// func RefreshToken(c *fiber.Ctx) error {
// 	headers := c.GetReqHeaders()
// 	return c.SendString(headers["Authorization"])
// }

func createClaims(user *models.User) jwt.MapClaims {
	return jwt.MapClaims{
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now()),
		"id":        user.ID,
		"phone":     user.Phone,
		"image":     user.Image,
		"email":     user.Email,
		"is_shop":   user.IsShop,
		"name":      user.Name,
	}
}
