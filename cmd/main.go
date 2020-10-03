package main

import (
	"github.com/domgolonka/kashoo-geo/internal/app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	_ "github.com/domgolonka/kashoo-geo/docs"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"net/http"
)

// @title Kashoo Echo
// @version 1.0
// @description This is a simple two-endpoint echo/golang project.

// @contact.name Dominik Golonka
// @contact.url http://www.github.com/domgolonka
// @contact.email domgolonka@gmail.com

// @schemes http
// @host 127.0.0.1:1323
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders:     []string{},
		AllowCredentials: false,
		ExposeHeaders:    []string{},
		MaxAge:           0,
	}))

	e.GET("/geolocate/:ip", handlers.GeoLocate)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
