package middleware

import "net/http"

type WrapperWritter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWritter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}
