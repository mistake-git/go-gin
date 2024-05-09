package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin/models" // Import the models package to access ContactForm
)

func ContactCreate(c *gin.Context) {
	var form models.ContactForm // Use ContactForm from the models package

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの処理に失敗しました"})
		return
	}

	contact := models.NewContact(&form)
	if err := contact.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := contact.SendConfirmMail(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "メールの送信に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "メールが正常に送信されました"})
}
