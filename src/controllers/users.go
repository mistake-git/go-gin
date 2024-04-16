package controllers

import (
	"net/http" 
	"github.com/gin-gonic/gin"
)

func UserLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"name": "Kudo",
		"language": "Baaroo",
	})
}