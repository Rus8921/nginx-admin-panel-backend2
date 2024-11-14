package NginxServer

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// InitTestDbNginx initializes an in-memory SQLite database for testing purposes.
// It creates a table for NginxServer and populates it with initial data.
// Returns the database connection and any error encountered.
func InitTestDbNginx() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&NginxServer{})
	if err != nil {
		return nil, err
	}

	servers := []NginxServer{
		{Ip: "test1@example.com", Domain: "testuser1", IsActive: true},
		{Ip: "test2@example.com", Domain: "testuser2", IsActive: false},
		{Ip: "test3@example.com", Domain: "testuser3", IsActive: true},
	}

	for i := range servers {
		if err := CreateNginxServer(db, &servers[i]); err != nil {
			return nil, err // Error handling for server creation
		}
	}
	return db, nil
}

// TestCreateNginxSeerver tests the CreateNginxServer function.
// It verifies that a new NginxServer can be created and retrieved correctly.
func TestCreateNginxSeerver(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxServer(db, server))

	var foundServer NginxServer

	assert.NoError(t, db.Where("id = ?", 4).First(&foundServer).Error)
	assert.Equal(t, "testIp", foundServer.Ip)
	assert.Equal(t, "testDomain", foundServer.Domain)
}

// TestGetNginxServer tests the GetNginxServer function.
// It verifies that NginxServers can be retrieved by their ID.
func TestGetNginxServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	servers, err := GetNginxServersAll(db)
	assert.NoError(t, err)
	assert.Len(t, servers, 3)

	foundServer1, err := GetNginxServer(db, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), foundServer1.ID)

	foundServer2, err := GetNginxServer(db, 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(2), foundServer2.ID)

	foundServer3, err := GetNginxServer(db, 3)
	assert.NoError(t, err)
	assert.Equal(t, uint(3), foundServer3.ID)
}

// TestGetNginxServersAll tests the GetNginxServersAll function.
// It verifies that all NginxServers can be retrieved.
func TestGetNginxServersAll(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	servers, err := GetNginxServersAll(db)
	assert.NoError(t, err)
	assert.Len(t, servers, 3)
}

// TestDeleteNginxServer tests the DeleteNginxServer function.
// It verifies that an NginxServer can be deleted and is no longer retrievable.
func TestDeleteNginxServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxServer(db, server))

	assert.NoError(t, DeleteNginxServer(db, 1))

	var foundServer NginxServer

	assert.Error(t, db.Where("id = ?", 1).First(&foundServer).Error)
}

// TestUpdateNginxServer tests the UpdateNginxServer function.
// It verifies that an NginxServer can be updated and the changes are saved correctly.
func TestUpdateNginxServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxServer(db, server))

	server.Ip = "newIp"
	server.Domain = "newDomain"

	updatedServer, err := UpdateNginxServer(db, 1, *server)
	assert.NoError(t, err)
	assert.Equal(t, "newIp", updatedServer.Ip)
	assert.Equal(t, "newDomain", updatedServer.Domain)
}

// ActivateOrUnactivateServer tests the ActivateOrUnactivateServer function.
// It verifies that an NginxServer can be activated and deactivated correctly.
func TestActivateOrUnactivateServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxServer(db, server))

	// Activate the server
	assert.NoError(t, ActivateOrUnactivateServer(db, server.ID))
	var foundServer NginxServer
	assert.NoError(t, db.Where("id = ?", server.ID).First(&foundServer).Error)
	assert.True(t, foundServer.IsActive)

	// Deactivate the server
	assert.NoError(t, ActivateOrUnactivateServer(db, server.ID))
	assert.NoError(t, db.Where("id = ?", server.ID).First(&foundServer).Error)
	assert.False(t, foundServer.IsActive)

	// Edge case: Non-existent server
	err = ActivateOrUnactivateServer(db, 999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error finding server")
}
