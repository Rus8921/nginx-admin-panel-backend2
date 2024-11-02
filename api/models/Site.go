package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	serverName string
	url        string
	isActive   bool
}
