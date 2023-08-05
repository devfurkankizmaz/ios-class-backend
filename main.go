/*
 * @File: main.go
 * @Author: Furkan Kızmaz (dev.furkan@outlook.com)
 */
package main

import (
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(middleware.Static("/docs"))
	app := configs.App()
	routes.Setup(app.DB, server)

	server.GET("/", HealthCheck)

	server.Logger.Fatal(server.Start("0.0.0.0:3000"))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
