package main

import (
	"log"

	"github.com/xydens/golang-blog-app/pkg/application"
	"github.com/xydens/golang-blog-app/pkg/exithandler"
)

func main() {

	_, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	exithandler.Init(func() {
	})
}
