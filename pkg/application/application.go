package application

import "fmt"

type Application struct {
}

func Get() (*Application, error) {
	fmt.Println("app start")
	return &Application{}, nil
}
