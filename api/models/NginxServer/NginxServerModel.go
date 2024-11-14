package NginxServer

import (
	"fmt"
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"

	//"golang.org/x/net/ipv4"
	"gorm.io/gorm"
)

// NginxServer represents an Nginx server with its associated properties.
type NginxServer struct {
	gorm.Model
	//Ip     ipv4.Conn `gorm:"unique;not null"` // потом надо переписать на него, если это возможно
	Ip            string        `gorm:"unique;not null"` // IP address of the server
	Domain        string        `gorm:"unique;not null"` // Domain name of the server
	ServerName    string        // Server name
	IsActive      bool          // Indicates if the server is active
	SitesOfServer []models.Site // List of sites associated with the server
}

// CreateNginxServer creates a new Nginx server record in the database.
// The server is initially set to inactive.
func CreateNginxServer(db *gorm.DB, server *NginxServer) error {
	server.IsActive = false
	result := db.Create(server)
	return result.Error
}

// GetNginxServer retrieves an Nginx server by its ID.
// Returns the server and an error if the server does not exist.
func GetNginxServer(db *gorm.DB, id uint) (NginxServer, error) {
	var server NginxServer
	if err := db.Where("id = ?", id).First(&server).Error; err != nil {
		return NginxServer{}, fmt.Errorf("server does not exist: %w", err)
	}
	return server, nil
}

// GetNginxServersAll retrieves all Nginx servers from the database.
// Returns a slice of servers and an error if any occurs.
func GetNginxServersAll(db *gorm.DB) ([]NginxServer, error) {
	var servers []NginxServer
	if err := db.Find(&servers).Error; err != nil {
		return nil, fmt.Errorf("error getting servers: %w", err)
	}
	return servers, nil
}

func GetAllSitesOfServer(db *gorm.DB, id uint) ([]models.Site, error) {
	var server NginxServer
	server, err := GetNginxServer(db, id)
	if err != nil {
		return nil, fmt.Errorf("error finding server: %w", err)
	}
	var sites []models.Site
	err = db.Model(&server).Association("SitesOfServer").Find(&sites)
	return sites, err
}

// DeleteNginxServer deletes an Nginx server by its ID.
// Returns an error if the server does not exist or if any other error occurs.
func DeleteNginxServer(db *gorm.DB, id uint) error {
	result := db.Delete(&NginxServer{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting server: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("server with Id %d not found", id)
	}
	return nil
}

// UpdateNginxServer updates an existing Nginx server with new data.
// Returns the updated server and an error if any occurs.
func UpdateNginxServer(db *gorm.DB, id uint, updatedServer NginxServer) (NginxServer, error) {
	var server NginxServer
	server, err := GetNginxServer(db, id)
	if err != nil {
		return NginxServer{}, fmt.Errorf("error finding server: %w", err)
	}
	if updatedServer.Ip != "" {
		server.Ip = updatedServer.Ip
	}
	if updatedServer.Domain != "" {
		server.Domain = updatedServer.Domain
	}
	if err := db.Save(&server).Error; err != nil {
		return NginxServer{}, fmt.Errorf("error saving server: %w", err)
	}
	return server, nil
}

// ActivateOrUnactivateServer toggles the active status of an Nginx server by its ID.
// Returns an error if the server does not exist or if any other error occurs.
func ActivateOrUnactivateServer(db *gorm.DB, id uint) error {
	var server NginxServer
	server, err := GetNginxServer(db, id)
	if err != nil {
		return fmt.Errorf("error finding server: %w", err)
	}
	if server.IsActive == true {
		server.IsActive = false
		if err := db.Save(&server).Error; err != nil {
			return fmt.Errorf("error saving server: %w", err)
		}
	} else {
		server.IsActive = true
		if err := db.Save(&server).Error; err != nil {
			return fmt.Errorf("error saving server: %w", err)
		}
	}
	return nil
}
