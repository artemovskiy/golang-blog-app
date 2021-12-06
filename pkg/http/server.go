package http

import "net/http"

func InitServer() error {
	return http.ListenAndServe("0.0.0.0:7777", CreateHandler())
}
