package handler

import (
	"github.com/Gopher4U/sturdy-fiesta/RESTapi/mqtt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"net/url"
)

func Health(c echo.Context) error {
	return c.JSON(http.StatusCreated, "OK")
}

func (conf *ConfigHandler) LigarLed(c echo.Context) error {
	uri, err := url.Parse(mqtt.BuildUri(conf.Config))
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

func (conf *ConfigHandler) DesligarLed(c echo.Context) error {
	uri, err := url.Parse(mqtt.BuildUri(conf.Config))
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