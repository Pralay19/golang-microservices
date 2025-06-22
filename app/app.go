package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartUp() {
	mapUrls()
	fmt.Println("The server is running...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
