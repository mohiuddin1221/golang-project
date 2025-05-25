package models

type UserProfile struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
