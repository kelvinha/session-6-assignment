package users

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service Service
}

func NewUserController(service Service) UserController {
	return UserController{
		service: service,
	}
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

func (ctrl *UserController) HalamanLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func (ctrl *UserController) HalamanHome(c echo.Context) error {
	session, _ := store.Get(c.Request(), SESSION_ID)
	if len(session.Values) == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"username": session.Values["username"],
	})
}

func (ctrl *UserController) Login(c echo.Context) (err error) {
	user := new(User)
	err = c.Bind(user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = ctrl.service.Login(user.Username, user.Password, c)
	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (ctrl *UserController) Logout(c echo.Context) (err error) {
	err = ctrl.service.Logout(c)
	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/home")
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
