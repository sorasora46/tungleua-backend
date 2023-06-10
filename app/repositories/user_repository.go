package repositories

import (
	"fmt"

	"github.com/sorasora46/Tungleua-backend/app/models"
	"github.com/sorasora46/Tungleua-backend/app/utils"
)

func GetUserById(userID string) (*models.User, error) {
	user := new(models.User)
	result := utils.DB.Find(user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	result := utils.DB.Find(user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByPhone(phone string) (*models.User, error) {
	user := new(models.User)
	result := utils.DB.Find(user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByName(name string) (*models.User, error) {
	user := new(models.User)
	result := utils.DB.Find(user, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func CheckIsUserExist(user *models.User) (bool, error) {
	result := utils.DB.Where("email = ?", user.Email).Or("name = ?", user.Name).Or("phone = ?", user.Phone).First(user)
	if result.Error != nil {
		return false, result.Error
	}

	fmt.Println(user)

	return user.ID != "", nil
}

// func CreateUser(user *models.User, user_password *models.Password) error {
func CreateUser(user *models.User) error {
	user_result := utils.DB.Create(user)
	if user_result.Error != nil {
		return user_result.Error
	}

	// pass_result := utils.DB.Create(user_password)
	// if pass_result.Error != nil {
	// 	return user_result.Error
	// }

	return nil
}
