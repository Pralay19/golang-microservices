package controllers

import (
	"fmt"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Page!</h1>"))
	fmt.Println("The server is running...")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
