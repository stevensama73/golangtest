package controllers

import (
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type PageController struct{
	htmlPath string
}

func (controller PageController) Index(c echo.Context) error {
	return c.File(filepath.Join(controller.htmlPath, "index.html"))
}

func (controller PageController) Input(c echo.Context) error {
	return c.File(filepath.Join(controller.htmlPath, "input.html"))
}

func (controller PageController) Output(c echo.Context) error {
	return c.File(filepath.Join(controller.htmlPath, "output.html"))
}

func NewPageController() *PageController {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	htmlPath := filepath.Join(dir, "templates")
	return &PageController{htmlPath: htmlPath}
}
