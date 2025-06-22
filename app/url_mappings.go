package app

import (
	"coursemicro/controllers"
)

func mapUrls() {
	router.GET("/", controllers.ServeHome)
	router.GET("/users", controllers.GetUser)
}
