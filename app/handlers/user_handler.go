package handlers

import (
	"fmt"
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

	result, err_result := repositories.UpdateUserById(userID, updates)
	if err_result != nil {
		return err_result
	}

	return c.SendString(result)
}

func GetUserById(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := repositories.GetUserById(userID)
	if err != nil {
		return err
	}

	fmt.Println(user.Image)

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

func CheckIsUserExist(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	isFound, err := repositories.CheckIsUserExist(req)
	if err != nil {
		return err
	}

	return c.JSON(map[string]bool{
		"isFound": isFound,
	})
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
