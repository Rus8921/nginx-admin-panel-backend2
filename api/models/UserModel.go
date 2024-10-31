package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	gorm.Model
	Id           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique;not null"` // Уникальный email
	Username     string `gorm:"unique;not null"` // Уникальное имя пользователя
	HashPassword string `gorm:"not null"`        // Хэшированный пароль
}

// CreateUser создает нового пользователя в базе данных
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func DeleteUser(db *gorm.DB, user *User) error {
	result := db.Delete(user)
	return result.Error
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashPassword), nil
}

func checkPassword(hashPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, fmt.Errorf("error checking password: %w", err)
	}
	return true, nil
}

func validateUser(db *gorm.DB, user User) error {
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		return fmt.Errorf("user already exist: %w", err)
	}
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		return fmt.Errorf("user email does not exist: %w", err)
	}
	return nil
}

func RegistrateUser(db *gorm.DB, user User) error {
	if err := validateUser(db, user); err != nil {
		return err
	}
	hashedPassword, err := hashPassword(user.HashPassword)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	user.HashPassword = hashedPassword

	if err := CreateUser(db, &user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}
