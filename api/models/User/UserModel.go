package User

import (
	"fmt"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	Email        string                  `gorm:"unique;not null"` // Unique email
	Username     string                  `gorm:"unique;not null"` // Unique username
	HashPassword string                  `gorm:"not null"`        // Hashed password
	Permissions  []Permission.Permission // One-to-one relationship with the Permission table
}

// CreateUser creates a new user in the database
// Parameters:
// - db: the database connection
// - user: the user to be created
// Returns:
// - error: an error if the creation fails
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

// hashPassword hashes a plain text password
// Parameters:
// - password: the plain text password
// Returns:
// - string: the hashed password
// - error: an error if the hashing fails
func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashPassword), nil
}

// CheckPassword checks if the provided password matches the hashed password
// Parameters:
// - hashPassword: the hashed password
// - password: the plain text password
// Returns:
// - bool: true if the passwords match, false otherwise
func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// validateUser checks if the username or email already exists in the database
// Parameters:
// - db: the database connection
// - user: the user to be validated
// Returns:
// - error: an error if the username or email already exists
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

// LoginUser logs in a user by checking the username and password
// Parameters:
// - db: the database connection
// - username: the username of the user
// - password: the plain text password of the user
// Returns:
// - *User: the logged-in user
// - error: an error if the login fails
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

// RegistrateUser registers a new user in the database
// Parameters:
// - db: the database connection
// - user: the user to be registered
// Returns:
// - error: an error if the registration fails
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

// GetUserById retrieves a user by their ID
// Parameters:
// - db: the database connection
// - id: the ID of the user
// Returns:
// - User: the retrieved user
// - error: an error if the retrieval fails
func GetUserById(db *gorm.DB, id uint) (User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("user does not finede: %w", err)
	}
	return user, nil
}

// GetUserByUsername retrieves a user by their username
// Parameters:
// - db: the database connection
// - username: the username of the user
// Returns:
// - User: the retrieved user
// - error: an error if the retrieval fails
func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, fmt.Errorf("user does not finede: %w", err)
	}
	return user, nil
}

// DeleteUser deletes a user by their ID
// Parameters:
// - db: the database connection
// - id: the ID of the user to be deleted
// Returns:
// - error: an error if the deletion fails
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

// UpdateUser updates a user's information
// Parameters:
// - db: the database connection
// - username: the username of the user to be updated
// - currentPassword: the current password of the user
// - updatedUser: the updated user information
// Returns:
// - User: the updated user
// - error: an error if the update fails
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
