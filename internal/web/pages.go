package web

import (
	"embed"
)

//go:embed *.html
var pages embed.FS
