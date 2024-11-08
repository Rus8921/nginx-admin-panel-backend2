package Admin

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type Admin struct {
	gorm.Model
	Email        string                  `gorm:"unique;not null"` // Уникальный email
	Username     string                  `gorm:"unique;not null"` // Уникальное имя пользователя
	HashPassword string                  `gorm:"not null"`        // Хэшированный пароль
	Permissions  []Permission.Permission `gorm:"foreignKey:ApproveAdminID"`
}

// CreateAdmin создает нового пользователя в базе данных
func CreateAdmin(db *gorm.DB, user *Admin) error {
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

func validateAdmin(db *gorm.DB, user Admin) error {
	var existingAdmin Admin
	if err := db.Where("username = ?", user.Username).First(&existingAdmin).Error; err == nil {
		return fmt.Errorf("username already exists")
	}
	if err := db.Where("email = ?", user.Email).First(&existingAdmin).Error; err == nil {
		return fmt.Errorf("email already exists")
	}
	return nil
}

func LoginAdmin(db *gorm.DB, username string, password string) (*Admin, error) {
	var admin Admin
	if err := db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, fmt.Errorf("user does not exist: %w", err)
	}
	result := CheckPassword(admin.HashPassword, password)
	if result == false {
		return nil, fmt.Errorf("invalid password")
	}
	return &admin, nil
}

func RegistrateAdmin(db *gorm.DB, user Admin) error {
	if err := validateAdmin(db, user); err != nil {
		return err
	}
	hashedPassword, err := hashPassword(user.HashPassword)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	user.HashPassword = hashedPassword

	if err := CreateAdmin(db, &user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func GetAdminById(db *gorm.DB, id uint) (Admin, error) {
	var admin Admin
	if err := db.Where("id = ?", id).First(&admin).Error; err != nil {
		return Admin{}, fmt.Errorf("user does not finede: %w", err)
	}
	return admin, nil
}

func DeleteAdmin(db *gorm.DB, id uint) error {
	result := db.Delete(&Admin{}, id)

	if result.Error != nil {
		return fmt.Errorf("ошибка удаления администратора: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("администратор с ID %d не найден", id)
	}
	return nil
}

func UpdateAdmin(db *gorm.DB, username, currentPassword string, updatedUser Admin) (Admin, error) {
	var admin Admin
	if err := db.Where("username = ?", username).First(&admin).Error; err != nil {
		return Admin{}, fmt.Errorf("admin does not exist: %w", err)
	}

	if !CheckPassword(admin.HashPassword, currentPassword) {
		return Admin{}, fmt.Errorf("invalid current password")
	}

	if updatedUser.Email != "" {
		admin.Email = updatedUser.Email
	}
	if updatedUser.Username != "" {
		admin.Username = updatedUser.Username
	}
	if updatedUser.HashPassword != "" {
		hashedPassword, err := hashPassword(updatedUser.HashPassword)
		if err != nil {
			return Admin{}, fmt.Errorf("admin does not exist: %w", err)
		}
		admin.HashPassword = hashedPassword
	}

	if err := db.Save(&admin).Error; err != nil {
		return Admin{}, fmt.Errorf("error updating admin: %w", err)
	}

	return admin, nil
}
