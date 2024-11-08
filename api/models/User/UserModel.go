package User

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null"` // Уникальный email
	Username     string `gorm:"unique;not null"` // Уникальное имя пользователя
	HashPassword string `gorm:"not null"`        // Хэшированный пароль
	Permissions  []Permission.Permission
}

// CreateUser создает нового пользователя в базе данных
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashPassword), nil
}

func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func validateUser(db *gorm.DB, user User) error {
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("username already exists")
	}
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("email already exists")
	}
	return nil
}

func LoginUser(db *gorm.DB, username string, password string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user does not exist: %w", err)
	}
	result := CheckPassword(user.HashPassword, password)
	if result == false {
		return nil, fmt.Errorf("invalid password")
	}
	return &user, nil
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

func GetUserById(db *gorm.DB, id uint) (User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("user does not finede: %w", err)
	}
	return user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("user does not finede: %w", err)
	}
	return user, nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&User{}, id)

	if result.Error != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("пользователь с ID %d не найден", id)
	}
	return nil
}

func UpdateUser(db *gorm.DB, username, currentPassword string, updatedUser User) (User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("user does not exist: %w", err)
	}

	if !CheckPassword(user.HashPassword, currentPassword) {
		return User{}, fmt.Errorf("invalid current password")
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}
	if updatedUser.Username != "" {
		user.Username = updatedUser.Username
	}
	if updatedUser.HashPassword != "" {
		hashedPassword, err := hashPassword(updatedUser.HashPassword)
		if err != nil {
			return User{}, fmt.Errorf("user does not exist: %w", err)
		}
		user.HashPassword = hashedPassword
	}

	if err := db.Save(&user).Error; err != nil {
		return User{}, fmt.Errorf("error updating user: %w", err)
	}

	return user, nil
}
