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
		hashedPassword, err := hashPassword("password" + string(rune(i+1)))
		if err != nil {
			return nil, err // Обработка ошибки хеширования
		}
		users[i].HashPassword = hashedPassword

		if err := CreateUser(db, &users[i]); err != nil {
			return nil, err // Обработка ошибки создания пользователя
		}
	}

	return db, nil
}

// прошел тестирование
func TestCheckPassword(t *testing.T) {
	password := "gtjfdniugjnrtri"
	hashedPassword, _ := hashPassword(password)
	assert.True(t, CheckPassword(hashedPassword, password)) // Проверка на верный пароль

	assert.False(t, CheckPassword(hashedPassword, "wrong")) // Проверка на неверный пароль
}

// прошел тестирование
func TestHashPassword(t *testing.T) {
	password := "dsfdsferdwe"
	hashedPassword, err := hashPassword(password)

	assert.NoError(t, err)             // Проверяем отсутствие ошибки при хешировании
	assert.NotEmpty(t, hashedPassword) // Проверяем что пароль не пустой

	assert.True(t, CheckPassword(hashedPassword, password)) // Проверка на верный пароль
}

// прошел тестирование
func TestCreateUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err)

	user := User{Email: "newuser@example.com", Username: "newuser"}
	hashedPassword, err := hashPassword("hashedpassword")
	assert.NoError(t, err) // Проверяем отсутствие ошибки при хешировании
	user.HashPassword = hashedPassword

	err = CreateUser(db, &user)
	assert.NoError(t, err) // Проверяем отсутствие ошибки при создании пользователя

	var foundUser User
	err = db.First(&foundUser, "email = ?", user.Email).Error
	assert.NoError(t, err)                             // Проверяем наличие ошибки при поиске пользователя
	assert.Equal(t, user.Username, foundUser.Username) // Убедимся что пользователь создан
}

// прошел тестирование
func TestValidateUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	user1 := &User{Email: "test@example.com", Username: "testuser"}
	hashedPassword, _ := hashPassword("hashedpassword")
	user1.HashPassword = hashedPassword

	err = CreateUser(db, user1)
	assert.NoError(t, err) // Проверяем отсутствие ошибки при создании пользователя

	user2 := User{Email: "test@example.com", Username: "testuser2"}
	err = validateUser(db, user2)
	assert.Error(t, err) // Ошибка при существующем email

	user3 := User{Email: "test2@example.com", Username: "testuser"}
	err = validateUser(db, user3)
	assert.Error(t, err) // Ошибка при существующем имени пользователя
}

// прошел тестирование
// TestLoginUser тестирует функцию логина.
func TestLoginUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	user := &User{Email: "login@example.com", Username: "loginuser"}
	hashedPassword, _ := hashPassword("password123")
	user.HashPassword = hashedPassword

	err = CreateUser(db, user)
	assert.NoError(t, err) // Проверяем отсутствие ошибки при создании пользователя

	loggedInUser, err := LoginUser(db, user.Username, "password123")
	assert.NoError(t, err)                                // Проверяем отсутствие ошибки при логине
	assert.Equal(t, user.Username, loggedInUser.Username) // Убедимся что пользователь залогинен

	_, err = LoginUser(db, user.Username+"wrong", "wrongpassword")
	assert.Error(t, err) // Проверка на неверные учетные данные
}

// прошел тестирование
// TestFindUser тестирует поиск пользователя.
func TestFindUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	foundUser1, err := GetUserByUsername(db, "testuser1")
	assert.NoError(t, err)                            // Проверяем отсутствие ошибки при поиске пользователя
	assert.Equal(t, foundUser1.Username, "testuser1") // Убедимся что пользователь найден

	foundUser2, err := GetUserByUsername(db, "nonexistentuser")
	assert.Error(t, err)                // Проверка на несуществующего пользователя
	assert.Equal(t, User{}, foundUser2) // Убедимся что пользователь не найден
}

// прошел тестирование
// TestDeleteUser тестирует удаление пользователя.
func TestDeleteUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	user := &User{Email: "delete@example.com", Username: "deleteuser", HashPassword: "hashedpassword"}
	assert.NoError(t, CreateUser(db, user)) // Проверяем отсутствие ошибки при создании пользователя

	assert.NoError(t, DeleteUser(db, user.ID)) // Проверяем отсутствие ошибки при удалении пользователя

	var foundUser User
	assert.Error(t, db.First(&foundUser, user.ID).Error) // Убедимся что пользователь удален
}

func TestGetUserById(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	foundUser1, err := GetUserById(db, 1)
	assert.NoError(t, err)                  // Проверяем отсутствие ошибки при поиске пользователя
	assert.Equal(t, uint(1), foundUser1.ID) // Убедимся что пользователь найден

	foundUser2, err := GetUserById(db, 999)
	assert.Error(t, err)                // Проверка на несуществующего пользователя
	assert.Equal(t, User{}, foundUser2) // Убедимся что пользователь не найден
}

func TestUpdateUser(t *testing.T) {
	db, err := InitTestDb()
	assert.NoError(t, err) // Проверяем отсутствие ошибки при инициализации базы данных

	user := &User{Email: "delete@example.com", Username: "deleteuser"}
	hashedPassword, err := hashPassword("hashedpassword")
	assert.NoError(t, err) // Проверяем отсутствие ошибки при хешировании
	user.HashPassword = hashedPassword
	assert.NoError(t, CreateUser(db, user)) // Прове��яем отсутствие ошибки при создании пользователя

	updatedUser, err := UpdateUser(db, user.Username, "hashedpassword", User{Username: "newusername", Email: "newemail"})
	assert.NoError(t, err)                               // Проверяем отсутствие ошибки при обновлении пользователя
	assert.Equal(t, updatedUser.Username, "newusername") // Убедимся что пользователь обновлен
	assert.Equal(t, updatedUser.Email, "newemail")       // Убедимся что пользователь обновлен
	hashedNewPassword, err := hashPassword("newpassword")
	assert.NoError(t, err)                                          // Проверяем отсутствие ошибки при хешировании нового пароля
	assert.True(t, CheckPassword(hashedNewPassword, "newpassword")) // Убедимся что новый пароль корректно хеширован

}
