package app

import (
	"coursemicro/controllers"
)

func mapUrls() {
	router.GET("/", controllers.ServeHome)
	router.GET("/users/:user_id", controllers.GetUser)
}
