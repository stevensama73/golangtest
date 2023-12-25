package controllers

import (
	"golangtest/src/models"
	"golangtest/src/services"
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProviderController struct {
	providerService services.ProviderService
}

func (controller ProviderController) AddProvider(c echo.Context) error {
	var provider models.Providers
	if err := c.Bind(&provider); err != nil {
		return err
	}
	
	result := controller.providerService.AddProvider(&provider)
	return c.JSON(http.StatusOK, result)
}

func (controller ProviderController) GetProviders(c echo.Context) error {
	result := controller.providerService.GetProviders()
	return c.JSON(http.StatusOK, result)
}

func (controller ProviderController) GeneratePhoneNumber(c echo.Context) error {
	result := controller.providerService.GeneratePhoneNumber()
	return c.JSON(http.StatusOK, result)
}

func NewProviderController(Conn *gorm.DB) ProviderController {
	service := services.NewProviderService(Conn)
	controller := ProviderController{
		providerService: service,
	}
	return controller
}