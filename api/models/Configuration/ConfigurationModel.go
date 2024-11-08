package Configuration

import (
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Site"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/api/models/Upstreams"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	parametrs string         `gorm:"not null"`
	Sites     []*models.Site `gorm:"many2many:configuration_site;"`
	Upstream  []Upstreams.Upstream
}
