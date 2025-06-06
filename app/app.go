package app

import (
	"coursemicro/controllers"
	"fmt"
	"log"
	"net/http"
)

func StartUp() {
	http.HandleFunc("/", controllers.ServeHome)
	http.HandleFunc("/users", controllers.GetUser)
	fmt.Println("The server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
