package main

import (
	"io"
	"session-6-assignment/database"
	"session-6-assignment/users"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	// Initializes database
	db := database.ConnectPGLocal()

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewUserController(userService)

	r.Renderer = NewRenderer("views/*.html", true)
	r.Any("/login", userHandler.HalamanLogin)
	r.Any("/home", userHandler.HalamanHome)
	r.POST("/login/store", userHandler.Login)
	r.POST("/logout", userHandler.Logout)

	r.Start(":9000")
}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}
