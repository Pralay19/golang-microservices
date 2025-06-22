package controllers

import (
	"coursemicro/services"
	"coursemicro/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServeHome(c *gin.Context) {
	// w.Write([]byte("<h1>Welcome to the Page!</h1>"))
	c.String(http.StatusOK, "<h1>Welcome to the Page!</h1>")
	fmt.Println("The server is running...")
}

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		// c.JSON(userErr.StatusCode, userErr)
		utils.RespondError(c, userErr)
		return
	}
	log.Printf("about to process %d", userId)
	// w.Write([]byte(fmt.Sprintf("<h1>This is the user-id: %s</h1>", userId)))

	// now we need to send this user-id to a service inorder to work with the database
	user, userErr := services.UsersService.GetUser(userId)
	if userErr != nil {
		// c.JSON(userErr.StatusCode, userErr)
		utils.RespondError(c, userErr)
		return
	}

	// c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}
