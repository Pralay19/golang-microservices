package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Page!</h1>"))
	fmt.Println("The server is running...")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userId := r.URL.Query().Get("id")
	log.Printf("about to process %s", userId)
	w.Write([]byte(fmt.Sprintf("<h1>This is the user-id: %s</h1>", userId)))
}
