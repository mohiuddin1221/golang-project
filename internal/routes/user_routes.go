package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohiuddin1221/golang-project/internal/services"
	"github.com/mohiuddin1221/golang-project/pkg/logger"
)

// UserRoutes defines the user-related routes
func UserRoutes(router *gin.Engine) {
	// Create user profile route
	router.POST("user/create", func(c *gin.Context) {
		var input services.CreateUserProfileInput

		// Bind JSON input to the struct
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the service to create a user profile
		UserProfile, err := services.CreateUserProfile(input)
		if err != nil {
			logger.Log.Error("Failed to create user profile:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
			return
		}

		// Log success and return response
		logger.Log.Info("User profile created successfully:", UserProfile)
		c.JSON(http.StatusOK, gin.H{"message": "User profile created successfully", "data": UserProfile})
	})
}
