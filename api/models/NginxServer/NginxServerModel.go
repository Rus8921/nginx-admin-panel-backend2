package NginxServer

import (
	"fmt"
	//"golang.org/x/net/ipv4"
	"gorm.io/gorm"
)

type NginxServer struct {
	gorm.Model
	//Ip     ipv4.Conn `gorm:"unique;not null"` // потом надо переписать на него, если это возможно
	Ip     string `gorm:"unique;not null"`
	Domain string `gorm:"unique;not null"`
}

func CreateNginxSeerver(db *gorm.DB, server *NginxServer) error {
	result := db.Create(server)
	return result.Error
}

func GetNginxServer(db *gorm.DB, id uint) (NginxServer, error) {
	var server NginxServer
	if err := db.Where("id = ?", id).First(&server).Error; err != nil {
		return NginxServer{}, fmt.Errorf("server does not exist: %w", err)
	}
	return server, nil
}

func GetNginxServersAll(db *gorm.DB) ([]NginxServer, error) {
	var servers []NginxServer
	if err := db.Find(&servers).Error; err != nil {
		return nil, fmt.Errorf("error getting servers: %w", err)
	}
	return servers, nil
}

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
