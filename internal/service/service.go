package service

import (
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/internal/webapp"
)

func Serve() {
	router := webapp.SetupRouter()
	router.Run(":8080")
}
