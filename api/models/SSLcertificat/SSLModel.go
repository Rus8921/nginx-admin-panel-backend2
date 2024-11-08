package SSLcertificat

import "gorm.io/gorm"

type SSL struct {
	gorm.Model
	FileCrt  string
	FileKey  string
	IsActive bool
	SiteID   uint
}
