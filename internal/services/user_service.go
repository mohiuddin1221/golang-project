package services

import (
	"github.com/google/uuid"
	"github.com/mohiuddin1221/golang-project/internal/database"
	"github.com/mohiuddin1221/golang-project/internal/models"
)

type CreateUserProfileInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func CreateuserProfile(input CreateUserProfileInput) (*models.UserProfile, error) {
	UserProfile := models.UserProfile{
		ID:      uuid.New().String(),
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Message: input.Message,
	}
	// if err := database.DB.Create(&UserProfile).Error; err != nil {
	// 	return nil, err
	// }
	// return &UserProfile, nil
	if err := database.DB.Create(&UserProfile).Error; err != nil {
		return nil, err
	}
	return &UserProfile, nil

}
