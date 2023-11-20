package main

import (
	"clean_go/internal/controllers"
	"clean_go/internal/database"
	"clean_go/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	// kafka := kafka.NewKafka()
	dbConn := database.SqliteConnection()

	//services
	productService := service.NewProductService(dbConn)

	// http controllers
	productHttp := controllers.NewProductHttpController(productService)
	productHttp.Routes(e)

	e.Logger.Fatal(e.Start(":4001"))
}
