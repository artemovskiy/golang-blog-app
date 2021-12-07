package http

import (
	"github.com/gorilla/mux"
	"github.com/xydens/golang-blog-app/pkg/core/post"
	"github.com/xydens/golang-blog-app/pkg/http/controller"
	"net/http"
)

func CreateHandler() (http.Handler, error) {
	var AppMux = Mux{Wrapped: mux.NewRouter()}

	PostService, err := post.NewPostService()
	if err != nil {
		return nil, err
	}

	postController := controller.PostController{
		PostService,
	}

	AppMux.HandleFunc("/post/{id}", postController.View).Methods("GET")
	AppMux.HandleFunc("/post", postController.Create).Methods("POST")

	return AppMux.Wrapped, nil
}
