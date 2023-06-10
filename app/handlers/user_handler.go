package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func GetUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := repositories.GetUserById(userID)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	id := uuid.New().String()

	// Get request body
	// req := new(models.UserWithPassword)
	// if err := c.BodyParser(&req); err != nil {
	// 	return err
	// }
	// req.ID = id
	req := new(models.User)
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	req.ID = id

	// Hashing password
	// user_pass := new(models.Password)
	// hashed_pass, hashed_err := bcrypt.GenerateFromPassword([]byte(req.Password), 3)
	// if hashed_err != nil {
	// 	return hashed_err
	// }
	// user_pass.UserID = id
	// user_pass.HashedPassword = string(hashed_pass)

	// create_err := repositories.CreateUser(&req.User, user_pass)
	create_err := repositories.CreateUser(req)
	if create_err != nil {
		return create_err
	}

	return nil
}
