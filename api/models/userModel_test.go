package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func InitTestDb() (*gorm.DB, error) {
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
		hashedPassword, err := hashPassword("password" + string(rune(i+1))) // Хешируем пароль для каждого пользователя
		if err != nil {
			return nil, err // Обработка ошибки хеширования
		}
		users[i].HashPassword = hashedPassword // Сохраняем хешированный пароль в структуру пользователя

		if err := CreateUser(db, &users[i]); err != nil {
			return nil, err // Обработка ошибки создания пользователя
		}
	}

	return db, nil
}

func TestCheckPassword(t *testing.T) {
	password := "gtjfdniugjnrtri"
	hashedPassword, _ := hashPassword(password)
	assert.True(t, CheckPassword(password, hashedPassword))

	assert.False(t, CheckPassword(hashedPassword, "wrong"))
}
func TestHashPassword(t *testing.T) {
	password := "dsfdsferdwe"
	hashedPassword, err := hashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	// Проверяем, что хешированный пароль соответствует оригинальному паролю
	assert.True(t, CheckPassword(hashedPassword, password))
}

func TestCreateUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	user := User{Email: "newuser@example.com", Username: "newuser"}
	hashedPassword, err := hashPassword("hashedpassword")
	assert.NoError(t, err)             // Проверяем отсутствие ошибки при хешировании
	user.HashPassword = hashedPassword // Используем хешированный пароль

	err = CreateUser(db, &user)
	assert.NoError(t, err)

	var foundUser User
	err = db.First(&foundUser, "email = ?", user.Email).Error
	assert.NoError(t, err)                             // Проверяем наличие ошибки при поиске пользователя
	assert.Equal(t, user.Username, foundUser.Username) // Убедимся что пользователь создан
}

func TestValidateUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	user1 := &User{Email: "test@example.com", Username: "testuser"}
	hashedPassword, _ := hashPassword("hashedpassword")
	user1.HashPassword = hashedPassword

	err = CreateUser(db, user1)
	assert.NoError(t, err)

	user2 := User{Email: "test@example.com", Username: "testuser2"}
	err = validateUser(db, user2)
	assert.Error(t, err) // Ошибка при существующем email

	user3 := User{Email: "test2@example.com", Username: "testuser"}
	err = validateUser(db, user3)
	assert.Error(t, err) // Ошибка при существующем имени пользователя
}

// TestLoginUser тестирует функцию логина.
func TestLoginUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	user := &User{Email: "login@example.com", Username: "loginuser"}
	hashedPassword, _ := hashPassword("password123")
	user.HashPassword = hashedPassword

	err = CreateUser(db, user)
	assert.NoError(t, err)

	loggedInUser, err := LoginUser(db, user.Username, "password123")
	assert.NoError(t, err)
	assert.Equal(t, user.Username, loggedInUser.Username)

	_, err = LoginUser(db, user.Username+"wrong", "wrongpassword")
	assert.Error(t, err) // Проверка на неверные учетные данные
}

// TestFindUser тестирует поиск пользователя.
func TestFindUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	foundUser1, err := GetUserByUsername(db, "testuser1")
	assert.NoError(t, err)
	assert.Equal(t, foundUser1.Username, "testuser1") // Убедимся что пользователь найден

	foundUser2, err := GetUserByUsername(db, "nonexistentuser")
	assert.Error(t, err) // Проверка на несуществующего пользователя
	assert.Nil(t, foundUser2)
}

// TestDeleteUser тестирует удаление пользователя.
func TestDeleteUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	user := &User{Email: "delete@example.com ", Username: "deleteuser ", HashPassword: "hashedpassword "}
	assert.NoError(t, CreateUser(db, user))

	assert.NoError(t, DeleteUser(db, user.Id))

	var foundUser User
	assert.Error(t, db.First(&foundUser, user.Id).Error) // Убедимся что пользователь удален
}
