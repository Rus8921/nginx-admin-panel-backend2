package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	ServerName     string
	Domain         string
	IsActive       bool
	NginxServerID  uint8
	Configurations []*models.Configurations `gorm:"many2many:configuration_site;"`
}
