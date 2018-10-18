package handler

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Gopher4U/sturdy-fiesta/RESTapi/mqtt"
	"github.com/labstack/echo"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusCreated, "OK")
}

func LigarLed(c echo.Context) error {
	uri, err := url.Parse(os.Getenv("MQTT_URI"))
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "led"
	}

	client := mqtt.Connect("pub", uri)
	client.Publish(topic, 0, false, "{\"energia\": true}")
	return c.JSON(http.StatusCreated, "OK")
}

func DesligarLed(c echo.Context) error {
	uri, err := url.Parse(os.Getenv("MQTT_URI"))
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "led"
	}

	client := mqtt.Connect("pub", uri)
	client.Publish(topic, 0, false, "{\"energia\": false}")
	return c.JSON(http.StatusCreated, "OK")
}