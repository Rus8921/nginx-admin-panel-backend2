package User

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// InitTestDbUser initializes an in-memory SQLite database for testing purposes.
// It creates a few test users and returns the database connection.
func InitTestDbUser() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	users := []User{
		{Email: "test1@example.com", Username: "testuser1"},
		{Email: "test2@example.com", Username: "testuser2"},
		{Email: "test3@example.com", Username: "testuser3"},
	}

	for i := range users {
		hashedPassword, err := hashPassword("password" + string(rune(i+1)))
		if err != nil {
			return nil, err // Error handling for password hashing
		}
		users[i].HashPassword = hashedPassword

		if err := CreateUser(db, &users[i]); err != nil {
			return nil, err // Error handling for user creation
		}
	}

	return db, nil
}

// TestCheckPassword tests the CheckPassword function.
func TestCheckPassword(t *testing.T) {
	password := "gtjfdniugjnrtri"
	hashedPassword, _ := hashPassword(password)
	assert.True(t, CheckPassword(hashedPassword, password)) // Check for correct password

	assert.False(t, CheckPassword(hashedPassword, "wrong")) // Check for incorrect password
}

// TestHashPassword tests the hashPassword function.
func TestHashPassword(t *testing.T) {
	password := "dsfdsferdwe"
	hashedPassword, err := hashPassword(password)

	assert.NoError(t, err)             // Check for no error during hashing
	assert.NotEmpty(t, hashedPassword) // Check that the hashed password is not empty

	assert.True(t, CheckPassword(hashedPassword, password)) // Check for correct password
}

// TestCreateUser tests the CreateUser function.
func TestCreateUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err)

	user := User{Email: "newuser@example.com", Username: "newuser"}
	hashedPassword, err := hashPassword("hashedpassword")
	assert.NoError(t, err) // Check for no error during hashing
	user.HashPassword = hashedPassword

	err = CreateUser(db, &user)
	assert.NoError(t, err) // Check for no error during user creation

	var foundUser User
	err = db.First(&foundUser, "email = ?", user.Email).Error
	assert.NoError(t, err)                             // Check for no error during user search
	assert.Equal(t, user.Username, foundUser.Username) // Ensure the user is created
}

// TestValidateUser tests the validateUser function.
func TestValidateUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	user1 := &User{Email: "test@example.com", Username: "testuser"}
	hashedPassword, _ := hashPassword("hashedpassword")
	user1.HashPassword = hashedPassword

	err = CreateUser(db, user1)
	assert.NoError(t, err) // Check for no error during user creation

	user2 := User{Email: "test@example.com", Username: "testuser2"}
	err = validateUser(db, user2)
	assert.Error(t, err) // Error for existing email

	user3 := User{Email: "test2@example.com", Username: "testuser"}
	err = validateUser(db, user3)
	assert.Error(t, err) // Error for existing username
}

// TestLoginUser tests the LoginUser function.
func TestLoginUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	user := &User{Email: "login@example.com", Username: "loginuser"}
	hashedPassword, _ := hashPassword("password123")
	user.HashPassword = hashedPassword

	err = CreateUser(db, user)
	assert.NoError(t, err) // Check for no error during user creation

	loggedInUser, err := LoginUser(db, user.Username, "password123")
	assert.NoError(t, err)                                // Check for no error during login
	assert.Equal(t, user.Username, loggedInUser.Username) // Ensure the user is logged in

	_, err = LoginUser(db, user.Username+"wrong", "wrongpassword")
	assert.Error(t, err) // Check for incorrect credentials
}

// TestFindUser tests the GetUserByUsername function.
func TestFindUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	foundUser1, err := GetUserByUsername(db, "testuser1")
	assert.NoError(t, err)                            // Check for no error during user search
	assert.Equal(t, foundUser1.Username, "testuser1") // Ensure the user is found

	foundUser2, err := GetUserByUsername(db, "nonexistentuser")
	assert.Error(t, err)                // Check for non-existent user
	assert.Equal(t, User{}, foundUser2) // Ensure the user is not found
}

// TestDeleteUser tests the DeleteUser function.
func TestDeleteUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	user := &User{Email: "delete@example.com", Username: "deleteuser", HashPassword: "hashedpassword"}
	assert.NoError(t, CreateUser(db, user)) // Check for no error during user creation

	assert.NoError(t, DeleteUser(db, user.ID)) // Check for no error during user deletion

	var foundUser User
	assert.Error(t, db.First(&foundUser, user.ID).Error) // Ensure the user is deleted
}

// TestGetUserById tests the GetUserById function.
func TestGetUserById(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	foundUser1, err := GetUserById(db, 1)
	assert.NoError(t, err)                  // Check for no error during user search
	assert.Equal(t, uint(1), foundUser1.ID) // Ensure the user is found

	foundUser2, err := GetUserById(db, 999)
	assert.Error(t, err)                // Check for non-existent user
	assert.Equal(t, User{}, foundUser2) // Ensure the user is not found
}

// TestUpdateUser tests the UpdateUser function.
func TestUpdateUser(t *testing.T) {
	db, err := InitTestDbUser()
	assert.NoError(t, err) // Check for no error during database initialization

	user := &User{Email: "delete@example.com", Username: "deleteuser"}
	hashedPassword, err := hashPassword("hashedpassword")
	assert.NoError(t, err) // Check for no error during hashing
	user.HashPassword = hashedPassword
	assert.NoError(t, CreateUser(db, user)) // Check for no error during user creation

	updatedUser, err := UpdateUser(db, user.Username, "hashedpassword", User{Username: "newusername", Email: "newemail"})
	assert.NoError(t, err)                               // Check for no error during user update
	assert.Equal(t, updatedUser.Username, "newusername") // Ensure the user is updated
	assert.Equal(t, updatedUser.Email, "newemail")       // Ensure the user is updated
	hashedNewPassword, err := hashPassword("newpassword")
	assert.NoError(t, err)                                          // Check for no error during new password hashing
	assert.True(t, CheckPassword(hashedNewPassword, "newpassword")) // Ensure the new password is correctly hashed
}
