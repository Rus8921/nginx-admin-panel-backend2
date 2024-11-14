package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func InitTestDbNginx() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Site{})
	if err != nil {
		return nil, err
	}

	servers := []Site{
		{SiteName: "test1@example.com", Domain: "testuser1", IsActive: true},
		{SiteName: "test2@example.com", Domain: "testuser2", IsActive: false},
		{SiteName: "test3@example.com", Domain: "testuser3", IsActive: true},
	}

	for i := range servers {
		if err := CreateSite(db, &servers[i]); err != nil {
			return nil, err // Обработка ошибки создания пользователя
		}
	}
	return db, nil
}

func TestCreateSite(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	site := &Site{SiteName: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateSite(db, site))

	var foundSite Site

	assert.NoError(t, db.Where("id = ?", 4).First(&foundSite).Error)
	assert.Equal(t, "testIp", foundSite.SiteName)
	assert.Equal(t, "testDomain", foundSite.Domain)
}

func TestGetSite(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	sites, err := GetSitesAll(db)
	assert.NoError(t, err)
	assert.Len(t, sites, 3)

	foundSite1, err := GetSite(db, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), foundSite1.ID)
}

func TestDeleteSite(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	err = DeleteSite(db, 1)
	assert.NoError(t, err)

	_, err = GetSite(db, 1)
	assert.Error(t, err)
}

func TestUpdateSite(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	updatedSite := Site{SiteName: "newSiteName", Domain: "newDomain"}
	site, err := UpdateSite(db, 1, updatedSite)
	assert.NoError(t, err)
	assert.Equal(t, "newSiteName", site.SiteName)
	assert.Equal(t, "newDomain", site.Domain)
}

//func TestActivateOrUnactivateSite(t *testing.T) {
//	db, err := InitTestDbNginx()
//	assert.NoError(t, err)
//
//	err = ActivateOrUnactivateSite(db, 1)
//	assert.NoError(t, err)
//
//	site, err := GetSite(db, 1)
//	assert.NoError(t, err)
//	assert.False(t, site.IsActive)
//}
