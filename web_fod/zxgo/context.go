package zxgo

import "net/http"

type Context struct {
	W http.ResponseWriter
	R *http.Request
}
