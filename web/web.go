package web

import (
	"embed"
)

//go:embed build/*
var WebFS embed.FS
