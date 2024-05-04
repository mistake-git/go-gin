package models

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/gomail.v2" // 正しいメールパッケージのインポート
)

type Contact struct {
	CompanyName       string `validate:"required,max=30"`
	DepartmentName    string `validate:"required,max=30"`
	PersonName        string `validate:"required,max=30"`
	PersonEmail       string `validate:"required,email,max=30"`
	PersonPhoneNumber string `validate:"required,len=10|len=11"`
	Subject           string `validate:"required,max=30"`
	Content           string `validate:"required"`
}

type ContactData interface {
	GetCompanyName() string
	GetDepartmentName() string
	GetPersonName() string
	GetPersonEmail() string
	GetPersonPhoneNumber() string
	GetSubject() string
	GetContent() string
}

func NewContact(data ContactData) *Contact {
	return &Contact{
			CompanyName:       data.GetCompanyName(),
			DepartmentName:    data.GetDepartmentName(),
			PersonName:        data.GetPersonName(),
			PersonEmail:       data.GetPersonEmail(),
			PersonPhoneNumber: data.GetPersonPhoneNumber(),
			Subject:           data.GetSubject(),
			Content:           data.GetContent(),
	}
}

func (c *Contact) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func (c *Contact) SendConfirmMail() error {
	m := gomail.NewMessage() // 正しい関数の呼び出し
	m.SetHeader("From", "acbmstk0402@gmail.com")
	m.SetHeader("To", c.PersonEmail)
	m.SetHeader("Subject", c.Subject)
	m.SetBody("text/plain", c.Content)

	d := gomail.NewDialer("smtp.gmail.com", 587, "quickly.t2024@gmail.com", "grpc lodj hhmg cura") // 正しい関数の呼び出し
	return d.DialAndSend(m)
}