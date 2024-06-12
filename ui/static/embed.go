/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/

package static

import "embed"

//go:embed *.css *.js icons/* *.map
var Files embed.FS
