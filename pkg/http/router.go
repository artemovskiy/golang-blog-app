package http

import (
	"github.com/gorilla/mux"
	"github.com/xydens/golang-blog-app/pkg/http/controller"
	"io"
	"net/http"
	"strings"
)

var postController = controller.PostController{}

func CreateHandler() http.Handler {
	var AppMux = mux.NewRouter()

	AppMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.Copy(writer, strings.NewReader("foo route response"))
	})

	AppMux.HandleFunc("/post/{id:[0-9]+}", postController.View).Methods("GET")
	AppMux.HandleFunc("/post", postController.Create).Methods("POST")

	return AppMux
}
