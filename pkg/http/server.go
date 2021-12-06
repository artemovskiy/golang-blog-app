package http

import "net/http"

func InitServer() error {
	handler, err := CreateHandler()
	if err != nil {
		return err
	}
	return http.ListenAndServe("0.0.0.0:7777", handler)
}
