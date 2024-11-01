package NginxServer

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

	err = db.AutoMigrate(&NginxServer{})
	if err != nil {
		return nil, err
	}

	servers := []NginxServer{
		{Ip: "test1@example.com", Domain: "testuser1"},
		{Ip: "test2@example.com", Domain: "testuser2"},
		{Ip: "test3@example.com", Domain: "testuser3"},
	}

	for i := range servers {
		if err := CreateNginxSeerver(db, &servers[i]); err != nil {
			return nil, err // Обработка ошибки создания пользователя
		}
	}
	return db, nil
}

func TestCreateNginxSeerver(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxSeerver(db, server))

	var foundServer NginxServer

	assert.NoError(t, db.Where("id = ?", 4).First(&foundServer).Error)
	assert.Equal(t, "testIp", foundServer.Ip)
	assert.Equal(t, "testDomain", foundServer.Domain)
}

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

func TestGetNginxServersAll(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	servers, err := GetNginxServersAll(db)
	assert.NoError(t, err)
	assert.Len(t, servers, 3)
}

func TestDeleteNginxServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxSeerver(db, server))

	assert.NoError(t, DeleteNginxServer(db, 1))

	var foundServer NginxServer

	assert.Error(t, db.Where("id = ?", 1).First(&foundServer).Error)
}

func TestUpdateNginxServer(t *testing.T) {
	db, err := InitTestDbNginx()
	assert.NoError(t, err)

	server := &NginxServer{Ip: "testIp", Domain: "testDomain"}
	assert.NoError(t, CreateNginxSeerver(db, server))

	server.Ip = "newIp"
	server.Domain = "newDomain"

	updatedServer, err := UpdateNginxServer(db, 1, *server)
	assert.NoError(t, err)
	assert.Equal(t, "newIp", updatedServer.Ip)
	assert.Equal(t, "newDomain", updatedServer.Domain)
}
