package models

import (
	"example.com/go-fiber-crud/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create inserts the user into the database
func (u *User) Create(db *gorm.DB) error {
	// Hash password first
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return db.Create(u).Error
}

func (u *User) Validate(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

// FindUserByEmail checks if a user exists by email
func FindUserByEmail(db *gorm.DB, email string, user *User) error {
	return db.Where("email = ?", email).First(user).Error
}
