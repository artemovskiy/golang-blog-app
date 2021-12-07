package http

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type Mux struct {
	Wrapped *mux.Router
}

func (m Mux) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request) error) *mux.Route {
	return m.Wrapped.HandleFunc(path, m.transformHandleFunc(f))
}

func (m Mux) transformHandleFunc(f func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			response := errorResponse{
				Message: err.Error(),
			}
			bytesResponse, err := json.Marshal(response)
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(500)
			io.Copy(w, bytes.NewReader(bytesResponse))
		}
	}
}
