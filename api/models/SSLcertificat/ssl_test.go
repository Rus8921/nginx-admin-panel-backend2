package SSLcertificat

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&SSL{})
	return db, nil
}

func TestCreateSSL(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl := &SSL{FileCrt: "path/to/cert.crt", FileKey: "path/to/key.key"}
	err = CreateSSL(db, ssl)
	assert.NoError(t, err)

	var result SSL
	db.First(&result, ssl.ID)
	assert.Equal(t, ssl.FileCrt, result.FileCrt)
	assert.Equal(t, ssl.FileKey, result.FileKey)
	assert.False(t, result.IsActive)
}

func TestGetSSL(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl := &SSL{FileCrt: "path/to/cert.crt", FileKey: "path/to/key.key"}
	db.Create(ssl)

	result, err := GetSSL(db, ssl.ID)
	assert.NoError(t, err)
	assert.Equal(t, ssl.FileCrt, result.FileCrt)
	assert.Equal(t, ssl.FileKey, result.FileKey)
}

func TestGetSSL_NotFound(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	_, err = GetSSL(db, 999)
	assert.Error(t, err)
}

func TestGetSSLCertificatesAll(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl1 := &SSL{FileCrt: "path/to/cert1.crt", FileKey: "path/to/key1.key"}
	ssl2 := &SSL{FileCrt: "path/to/cert2.crt", FileKey: "path/to/key2.key"}
	db.Create(ssl1)
	db.Create(ssl2)

	results, err := GetSSLCertificatesAll(db)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestDeleteSSL(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl := &SSL{FileCrt: "path/to/cert.crt", FileKey: "path/to/key.key"}
	db.Create(ssl)

	err = DeleteSSL(db, ssl.ID)
	assert.NoError(t, err)

	var result SSL
	err = db.First(&result, ssl.ID).Error
	assert.Error(t, err)
}

func TestDeleteSSL_NotFound(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	err = DeleteSSL(db, 999)
	assert.Error(t, err)
}

func TestUpdateSSL(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl := &SSL{FileCrt: "path/to/cert.crt", FileKey: "path/to/key.key"}
	db.Create(ssl)

	updatedSSL := SSL{FileCrt: "new/path/to/cert.crt"}
	result, err := UpdateSSL(db, ssl.ID, updatedSSL)
	assert.NoError(t, err)
	assert.Equal(t, updatedSSL.FileCrt, result.FileCrt)
}

func TestActivateOrUnactivateSSL(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	ssl := &SSL{FileCrt: "path/to/cert.crt", FileKey: "path/to/key.key"}
	db.Create(ssl)

	err = ActivateOrUnactivateSSL(db, ssl.ID)
	assert.NoError(t, err)

	var result SSL
	db.First(&result, ssl.ID)
	assert.True(t, result.IsActive)

	err = ActivateOrUnactivateSSL(db, ssl.ID)
	assert.NoError(t, err)

	db.First(&result, ssl.ID)
	assert.False(t, result.IsActive)
}
