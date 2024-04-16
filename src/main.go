package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/controllers"
)

func main() {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/users", controllers.UserLists)
	
	router.Run(":8000")
}
