package repositories

import (
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
