package users

import "html/template"

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"passwoed"`
}

type UserRegister struct {
	Username  string `form:"username" json:"username"`
	FirstName string `form:"first_name" json:"first_name"`
	LastName  string `form:"last_name" json:"last_name"`
	Password  string `form:"password" json:"passwoed"`
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}
