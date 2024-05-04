package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-gin/models"
)

type ContactForm struct {
	CompanyName       string `form:"company_name"`
	DepartmentName    string `form:"department_name"`
	PersonName        string `form:"person_name"`
	PersonEmail       string `form:"person_email"`
	PersonPhoneNumber string `form:"person_phone_number"`
	Subject           string `form:"subject"`
	Content           string `form:"content"`
}

// Implementing the ContactData interface
func (cf *ContactForm) GetCompanyName() string {
	return cf.CompanyName
}

func (cf *ContactForm) GetDepartmentName() string {
	return cf.DepartmentName
}

func (cf *ContactForm) GetPersonName() string {
	return cf.PersonName
}

func (cf *ContactForm) GetPersonEmail() string {
	return cf.PersonEmail
}

func (cf *ContactForm) GetPersonPhoneNumber() string {
	return cf.PersonPhoneNumber
}

func (cf *ContactForm) GetSubject() string {
	return cf.Subject
}

func (cf *ContactForm) GetContent() string {
	return cf.Content
}

func ContactCreate(c *gin.Context) {
	var form ContactForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "メールの送信に失敗しました"})
		return
	}

	contact := models.NewContact(&form)
	if err := contact.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := contact.SendConfirmMail(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "メールが送信されました"})
}
