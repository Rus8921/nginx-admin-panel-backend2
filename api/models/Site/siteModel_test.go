package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// InitTestDbNginx initializes an in-memory SQLite database for testing purposes.
// It creates a table for the Site model and populates it with initial data.
// Returns the database connection and any error encountered.
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
			return nil, err // Error handling for site creation
		}
	}
	return db, nil
}

// TestCreateSite tests the CreateSite function.
// It initializes the test database, creates a new site, and verifies that the site was created correctly.
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

// TestGetSite tests the GetSite function.
// It initializes the test database, retrieves all sites, and verifies that the correct number of sites is returned.
// It also retrieves a specific site by ID and verifies its properties.
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

// TestDeleteSite tests the DeleteSite function.
// It initializes the test database, deletes a site by ID, and verifies that the site was deleted correctly.
func TestDeleteSite(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	err = DeleteSite(db, 1)
	assert.NoError(t, err)

	_, err = GetSite(db, 1)
	assert.Error(t, err)
}

// TestUpdateSite tests the UpdateSite function.
// It initializes the test database, updates a site's properties, and verifies that the site was updated correctly.
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
//	server := &Site{SiteName: "testIp", Domain: "testDomain"}
//	assert.NoError(t, CreateSite(db, server))
//
//	// Activate the server
//	assert.NoError(t, ActivateOrUnactivateSite(db, server.ID))
//	var foundSite Site
//	assert.NoError(t, db.Where("id = ?", server.ID).First(&Site{}).Error)
//	assert.True(t, foundSite.IsActive)
//
//	// Deactivate the server
//	assert.NoError(t, ActivateOrUnactivateSite(db, server.ID))
//	assert.NoError(t, db.Where("id = ?", server.ID).First(&foundSite).Error)
//	assert.False(t, foundSite.IsActive)
//
//	// Edge case: Non-existent server
//	err = ActivateOrUnactivateSite(db, 999)
//	assert.Error(t, err)
//	assert.Contains(t, err.Error(), "error finding server")
//}
