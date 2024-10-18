package main

import (
	"fmt"

	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/internal/service"
)

func main() {
	fmt.Println("Starting service...")
	service.Serve()
}
