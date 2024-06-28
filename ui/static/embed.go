package static

import "embed"

//go:embed *.css *.js icons/*
var Files embed.FS
