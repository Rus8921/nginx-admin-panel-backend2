package models

import (
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
