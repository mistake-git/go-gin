package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/users", controllers.UserLists)
	apiV1.GET("/contacts", controllers.ContactCreate)
	
	router.Run(":8000")
}
