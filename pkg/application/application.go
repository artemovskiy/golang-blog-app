package application

import (
	"fmt"
	"github.com/xydens/golang-blog-app/pkg/http"
)

type Application struct {
}

func Get() (*Application, error) {
	fmt.Println("app start")
	err := http.InitServer()
	if err != nil {
		return nil, err
	}
	return &Application{}, nil
}
