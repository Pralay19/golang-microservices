package app

import (
	"coursemicro/controllers"
	"fmt"
	"log"
	"net/http"
)

func StartUp() {
	http.HandleFunc("/", controllers.ServeHome)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("The server is running...")
}
