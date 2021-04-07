package utils

import (
	"net/http"
	"net/http/httptest"
)

type IRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
	ADDVERSION(uri string)
	SERVEHTTP(rr *httptest.ResponseRecorder, req *http.Request)
}
