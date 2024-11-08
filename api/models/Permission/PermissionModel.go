package Permission

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	ApproveAdminID uint
	UserID         uint
	SiteID         uint
}
