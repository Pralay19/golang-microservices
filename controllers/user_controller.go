package controllers

import (
	"coursemicro/services"
	"coursemicro/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Page!</h1>"))
	fmt.Println("The server is running...")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		// log.Fatal(err)
		// w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("user-id must be a number"))
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonVal, _ := json.Marshal(userErr)
		w.WriteHeader(userErr.StatusCode)
		w.Write(jsonVal)
		return
	}
	log.Printf("about to process %d", userId)
	// w.Write([]byte(fmt.Sprintf("<h1>This is the user-id: %s</h1>", userId)))

	// now we need to send this user-id to a service inorder to work with the database
	user, userErr := services.UsersService.GetUser(userId)
	if userErr != nil {
		// log.Fatal(err)
		jsonVal, _ := json.Marshal(userErr)
		w.WriteHeader(userErr.StatusCode)
		w.Write(jsonVal)
		// w.WriteHeader(userErr.StatusCode)
		// w.Write(jsonVal)
		return
	}
	// w.Write([]byte(fmt.Sprintf("<h1>User ID: %d, First-Name: %s</h1>", user.Id, user.FirstName)))

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user) //this is used to return in a proper json format

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
