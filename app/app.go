package app

import (
	"coursemicro/controllers"
	"net/http"
)

func StartUp() {
	http.HandleFunc("/users", controllers.GetUser)
}
