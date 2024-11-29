package Admin

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Admin represents the user model
type Admin struct {
	gorm.Model
	Email        string                  `gorm:"unique;not null"`           // Unique email
	Username     string                  `gorm:"unique;not null"`           // Unique username
	HashPassword string                  `gorm:"not null"`                  // Hashed password
	Permissions  []Permission.Permission `gorm:"foreignKey:ApproveAdminID"` // One-to-one relationship with the Permission table
}

// CreateAdmin creates a new user in the database
// db: Database connection
// user: Admin user to be created
// Returns an error if the creation fails
func CreateAdmin(db *gorm.DB, user *Admin) error {
	result := db.Create(user)
	return result.Error
}

// hashPassword hashes the given password
// password: Plain text password
// Returns the hashed password or an error if hashing fails
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashPassword), nil
}

// CheckPassword compares a hashed password with a plain text password
// hashPassword: Hashed password
// password: Plain text password
// Returns true if the passwords match, false otherwise
func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// validateAdmin checks if the username or email already exists in the database
// db: Database connection
// user: Admin user to be validated
// Returns an error if the username or email already exists
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

// LoginAdmin authenticates an admin user
// db: Database connection
// username: Admin username
// password: Admin password
// Returns the authenticated admin user or an error if authentication fails
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

// RegistrateAdmin registers a new admin user
// db: Database connection
// user: Admin user to be registered
// Returns an error if registration fails
func RegistrateAdmin(db *gorm.DB, user Admin) error {
	if err := validateAdmin(db, user); err != nil {
		return err
	}
	hashedPassword, err := HashPassword(user.HashPassword)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}
	user.HashPassword = hashedPassword

	if err := CreateAdmin(db, &user); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

// GetAdminById retrieves an admin user by ID
// db: Database connection
// id: Admin user ID
// Returns the admin user or an error if retrieval fails
func GetAdminById(db *gorm.DB, id uint) (Admin, error) {
	var admin Admin
	if err := db.Where("id = ?", id).First(&admin).Error; err != nil {
		return Admin{}, fmt.Errorf("user does not exist: %w", err)
	}
	return admin, nil
}

// DeleteAdmin deletes an admin user by ID
// db: Database connection
// id: Admin user ID
// Returns an error if deletion fails or if the user is not found
func DeleteAdmin(db *gorm.DB, id uint) error {
	result := db.Delete(&Admin{}, id)

	if result.Error != nil {
		return fmt.Errorf("error deleting admin: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("admin with ID %d not found", id)
	}
	return nil
}

// UpdateAdmin updates an admin user's details
// db: Database connection
// username: Admin username
// currentPassword: Current password of the admin user
// updatedUser: Admin user with updated details
// Returns the updated admin user or an error if the update fails
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
		hashedPassword, err := HashPassword(updatedUser.HashPassword)
		if err != nil {
			return Admin{}, fmt.Errorf("error hashing password: %w", err)
		}
		admin.HashPassword = hashedPassword
	}

	if err := db.Save(&admin).Error; err != nil {
		return Admin{}, fmt.Errorf("error updating admin: %w", err)
	}

	return admin, nil
}
