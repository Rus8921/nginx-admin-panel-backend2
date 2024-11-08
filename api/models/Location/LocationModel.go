package Location

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	SiteID uint
}
