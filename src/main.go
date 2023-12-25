package main

import (
	"golangtest/framework/databases"
	"golangtest/src/controllers"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := databases.InitMysql()

	route := echo.New()
	store := sessions.NewCookieStore([]byte(os.Getenv("GOOGLEKEY")))
	route.Use(sessionMiddleware(store))
	route.Use(middleware.CORS())

	providerRoute := route.Group("api/provider")
	googleRoute := route.Group("api/google")

	pageController := controllers.NewPageController()
	providerController := controllers.NewProviderController(db)
	googleController := controllers.NewGoogleController()

	route.GET("/", pageController.Index)
	route.GET("/input", pageController.Input)
	route.GET("/output", pageController.Output)

	providerRoute.POST("/add", providerController.AddProvider)
	providerRoute.GET("/", providerController.GetProviders)
	providerRoute.GET("/generate/phoneNumber", providerController.GeneratePhoneNumber)

	googleRoute.GET("/callback", googleController.SignInCallback)

	route.Start(":8080")
}

func sessionMiddleware(store *sessions.CookieStore) echo.MiddlewareFunc {
	return session.MiddlewareWithConfig(session.Config{
		Store: store,
	})
}