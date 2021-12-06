package http

import (
	"github.com/gorilla/mux"
	"github.com/xydens/golang-blog-app/pkg/core/post"
	"github.com/xydens/golang-blog-app/pkg/http/controller"
	"io"
	"net/http"
	"strings"
)

func CreateHandler() (http.Handler, error) {
	var AppMux = mux.NewRouter()

	PostService, err := post.NewPostService()
	if err != nil {
		return nil, err
	}

	postController := controller.PostController{
		PostService,
	}

	AppMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.Copy(writer, strings.NewReader("foo route response"))
	})

	AppMux.HandleFunc("/post/{id}", postController.View).Methods("GET")
	AppMux.HandleFunc("/post", postController.Create).Methods("POST")

	return AppMux, nil
}
