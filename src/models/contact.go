package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/gomail.v2"
)

// ContactForm now part of the models package.
type ContactForm struct {
	CompanyName       string `form:"company_name"`
	DepartmentName    string `form:"department_name"`
	PersonName        string `form:"person_name"`
	PersonEmail       string `form:"person_email"`
	PersonPhoneNumber string `form:"person_phone_number"`
	Subject           string `form:"subject"`
	Content           string `form:"content"`
}

type Contact struct {
	CompanyName       string `validate:"required,max=30"`
	DepartmentName    string `validate:"required,max=30"`
	PersonName        string `validate:"required,max=30"`
	PersonEmail       string `validate:"required,email,max=30"`
	PersonPhoneNumber string `validate:"required,len=10|len=11"`
	Subject           string `validate:"required,max=30"`
	Content           string `validate:"required"`
}

func NewContact(data *ContactForm) *Contact {
	return &Contact{
		CompanyName:       data.CompanyName,
		DepartmentName:    data.DepartmentName,
		PersonName:        data.PersonName,
		PersonEmail:       data.PersonEmail,
		PersonPhoneNumber: data.PersonPhoneNumber,
		Subject:           data.Subject,
		Content:           data.Content,
	}
}

func (c *Contact) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Contact) SendConfirmMail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "quickly.t2024@gmail.com")
	m.SetHeader("To", c.PersonEmail)
	m.SetHeader("Subject", c.Subject)
	m.SetBody("text/plain", c.Content)

	d := gomail.NewDialer("smtp.gmail.com", 587, "quickly.t2024@gmail.com", "grpc lodj hhmg cura")
	return d.DialAndSend(m)
}
