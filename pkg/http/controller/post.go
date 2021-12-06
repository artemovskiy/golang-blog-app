package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xydens/golang-blog-app/pkg/core/post"
	"io"
	"log"
	"net/http"
)

type PostController struct {
	*post.PostService
}

func (c *PostController) View(writer http.ResponseWriter, request *http.Request) {
	pathParams := mux.Vars(request)

	id := pathParams["id"]

	post, err := c.PostService.Get(id)
	if err != nil {
		log.Fatal(err)
	}

	bytesResponse, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(writer, bytes.NewReader(bytesResponse))
}

func (c *PostController) Create(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var createData post.CreatePostData
	err := decoder.Decode(&createData)
	if err != nil {
		log.Fatal(err)
	}

	data, err := c.PostService.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}

	bytesResponse, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(writer, bytes.NewReader(bytesResponse))
}
