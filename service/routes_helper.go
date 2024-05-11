package service

import (
	"net/http"
	"path"
	"strings"
)

func setURL(r *http.Request, val string) {
	r.URL.Path = val
	r.URL.RawPath = val
}

func shiftURL(r *http.Request) string {
	head, tail := shiftPath(r.URL.Path)
	setURL(r, tail)
	return head
}

// p is an absolute path
func shiftPath(p string) (head string, tail string) {
	if p[0] != '/' {
		panic("service.shiftPath: p must be an absolute path")
	}
	p = path.Clean(p[1:])
	si := strings.Index(p, "/")
	if si == -1 {
		return p, "/"
	}
	return p[:si], p[si:]
}
