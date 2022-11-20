package users

import "html/template"

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"passwoed"`
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}
