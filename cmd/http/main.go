package main

import (
	"clean_go/internal/controllers"
	"clean_go/internal/database"
	"clean_go/internal/service"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

var (
	upgrader = websocket.Upgrader{}
)

func WebsocketHandler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			return
		}
	}(ws)

	for {
		// write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		fmt.Printf("message from client: %s\n", msg)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ws", WebsocketHandler)

	e.GET("/chat", WebsocketHandler)

	// kafka := kafka.NewKafka()
	dbConn := database.SqliteConnection()

	//services
	productService := service.NewProductService(dbConn)

	// http controllers
	productHttp := controllers.NewProductHttpController(productService)
	productHttp.Routes(e)

	// Get the port number from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Start the server on the specified port
	e.Logger.Fatal(e.Start(":" + port))
}
