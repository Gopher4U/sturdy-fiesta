package router

import (
	"github.com/Gopher4U/sturdy-fiesta/RESTapi/config"
	"github.com/labstack/echo"
	h "github.com/Gopher4U/sturdy-fiesta/RESTapi/handler"
)

func Routes(e *echo.Echo, c *config.Config) {

	handler := h.StartConfig(c)

	e.GET("/health", h.Health)
	e.POST("/luz/ligar", handler.LigarLed)
	e.POST("/luz/desligar", handler.DesligarLed)
}
