package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

type GoogleController struct{
	htmlPath string
}

func (controller GoogleController) SignIn(c echo.Context) error {
	return c.File(filepath.Join(controller.htmlPath, "index.html"))
}

func (controller GoogleController) SignInCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func NewGoogleController() *GoogleController {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	htmlPath := filepath.Join(dir, "templates")
	return &GoogleController{htmlPath: htmlPath}
}
