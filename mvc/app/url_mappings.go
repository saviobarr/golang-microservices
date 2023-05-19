package app

import "github.com/saviobarr/golang-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
