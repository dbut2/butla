package web

import (
	"embed"
	_ "embed"
)

//go:embed *.html
var pages embed.FS
