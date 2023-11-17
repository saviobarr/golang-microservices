package main

import (
	"fmt"

	"github.com/saviobarr/golang-microservices/mvc/app"
)

func main() {
	app.StartApp()
	defer fmt.Println("World")
	fmt.Println("Hello")

	defer func() {
		recover()
		fmt.Println("panic recovered")
	}()

	panic("panic")

}
