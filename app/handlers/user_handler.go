package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/repositories"
)

func UpdateUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	updates := make(map[string]interface{})
	if err := c.BodyParser(&updates); err != nil {
		return err
	}

	err := repositories.UpdateUserById(userID, updates)
	if err != nil {
		return err
	}

	return nil
}

func GetUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := repositories.GetUserById(userID)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"id":      user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"image":   string(user.Image),
		"is_shop": user.IsShop,
		"phone":   user.Phone,
	})
}

func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")

	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"id":      user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"image":   user.Image,
		"is_shop": user.IsShop,
		"phone":   user.Phone,
	})
}

func CheckDuplicateUser(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := repositories.CheckDuplicateUser(req)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(c *fiber.Ctx) error {
	// id := uuid.New().String()

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
	req.Name = strings.ToLower(req.Name)
	req.Email = strings.ToLower(req.Email)
	// req.ID = id

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
