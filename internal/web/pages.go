package web

import (
	_ "embed"
)

//go:embed html/index.html
var index []byte

//go:embed html/404.html
var e404 []byte

//go:embed html/500.html
var e500 []byte
