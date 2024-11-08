package Upstreams

import "gorm.io/gorm"

type Upstream struct {
	gorm.Model
	UpstreamID uint
	Priority   uint   `gorm:"not nul"`
	Parametr   string `gorm:"not nul"`
}
