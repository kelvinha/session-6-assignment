package users

import (
	"log"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

var SESSION_ID = "my-session-id"

func newCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}

func newPostgresStore() *pgstore.PGStore {
	url := "postgres://postgres:kelvin@127.0.0.1:5432/hactive8?sslmode=disable"
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
	if err != nil {
		log.Println("ERROR", err.Error())
		os.Exit(0)
	}

	return store
}

var store = newPostgresStore()

// var store = newCookieStore()

func (service *Service) Login(username, password string, c echo.Context) (err error) {
	var (
		getUser User
		request User
	)

	request.Username = username
	request.Password = password

	getUser, err = service.repository.Login(request)
	if err != nil {
		return
	}

	session, err := store.Get(c.Request(), SESSION_ID)
	if err != nil {
		return
	}

	session.Values["username"] = getUser.Username
	session.Save(c.Request(), c.Response())

	return
}

func (service *Service) Logout(c echo.Context) (err error) {
	session, _ := store.Get(c.Request(), SESSION_ID)
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return nil
}
