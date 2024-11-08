package Location

import (
	models "gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Site []*models.Site `gorm:"many2many:location_site;"`
}
