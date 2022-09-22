package web

import (
	_ "embed"
)

//go:embed index.html
var index []byte

//go:embed 404.html
var e404 []byte

//go:embed 500.html
var e500 []byte
