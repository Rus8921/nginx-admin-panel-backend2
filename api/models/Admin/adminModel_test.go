package Admin

import (
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Permission"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&Admin{}, &Permission.Permission{})
	return db
}

func TestCreateAdmin_Success(t *testing.T) {
	db := setupTestDB()
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: "password"}

	err := CreateAdmin(db, &admin)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCreateAdmin_DuplicateEmail(t *testing.T) {
	db := setupTestDB()
	admin1 := Admin{Email: "test@example.com", Username: "testuser1", HashPassword: "password"}
	admin2 := Admin{Email: "test@example.com", Username: "testuser2", HashPassword: "password"}

	CreateAdmin(db, &admin1)
	err := CreateAdmin(db, &admin2)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestHashPassword_Success(t *testing.T) {
	password := "password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		t.Fatalf("expected passwords to match, got %v", err)
	}
}

func TestCheckPassword_Success(t *testing.T) {
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	result := CheckPassword(string(hashedPassword), password)
	if !result {
		t.Fatalf("expected passwords to match, got false")
	}
}

func TestCheckPassword_Failure(t *testing.T) {
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	result := CheckPassword(string(hashedPassword), "wrongpassword")
	if result {
		t.Fatalf("expected passwords not to match, got true")
	}
}

func TestLoginAdmin_Success(t *testing.T) {
	db := setupTestDB()
	password := "password"
	hashedPassword, _ := HashPassword(password)
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: hashedPassword}
	CreateAdmin(db, &admin)

	loggedInAdmin, err := LoginAdmin(db, "testuser", password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if loggedInAdmin.Username != "testuser" {
		t.Fatalf("expected username to be 'testuser', got %v", loggedInAdmin.Username)
	}
}

func TestLoginAdmin_InvalidPassword(t *testing.T) {
	db := setupTestDB()
	password := "password"
	hashedPassword, _ := HashPassword(password)
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: hashedPassword}
	CreateAdmin(db, &admin)

	_, err := LoginAdmin(db, "testuser", "wrongpassword")
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestRegistrateAdmin_Success(t *testing.T) {
	db := setupTestDB()
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: "password"}

	err := RegistrateAdmin(db, admin)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRegistrateAdmin_DuplicateUsername(t *testing.T) {
	db := setupTestDB()
	admin1 := Admin{Email: "test1@example.com", Username: "testuser", HashPassword: "password"}
	admin2 := Admin{Email: "test2@example.com", Username: "testuser", HashPassword: "password"}

	RegistrateAdmin(db, admin1)
	err := RegistrateAdmin(db, admin2)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestGetAdminById_Success(t *testing.T) {
	db := setupTestDB()
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: "password"}
	CreateAdmin(db, &admin)

	retrievedAdmin, err := GetAdminById(db, admin.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if retrievedAdmin.Username != "testuser" {
		t.Fatalf("expected username to be 'testuser', got %v", retrievedAdmin.Username)
	}
}

func TestDeleteAdmin_Success(t *testing.T) {
	db := setupTestDB()
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: "password"}
	CreateAdmin(db, &admin)

	err := DeleteAdmin(db, admin.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUpdateAdmin_Success(t *testing.T) {
	db := setupTestDB()
	password := "password"
	hashedPassword, _ := HashPassword(password)
	admin := Admin{Email: "test@example.com", Username: "testuser", HashPassword: hashedPassword}
	CreateAdmin(db, &admin)

	updatedAdmin := Admin{Email: "new@example.com", Username: "newuser", HashPassword: "newpassword"}
	updatedAdmin, err := UpdateAdmin(db, "testuser", password, updatedAdmin)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updatedAdmin.Email != "new@example.com" {
		t.Fatalf("expected email to be 'new@example.com', got %v", updatedAdmin.Email)
	}
}
