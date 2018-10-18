package main

import (
	"github.com/Gopher4U/sturdy-fiesta/RESTapi/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	//uri, err := url.Parse(os.Getenv("MQTT_URI"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//topic := uri.Path[1:len(uri.Path)]
	//if topic == "" {
	//	topic = "led"
	//}
	//
	//go mqtt.Listen(uri, topic)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	router.Routes(e)
	e.Logger.Fatal(e.Start(":3000"))
}