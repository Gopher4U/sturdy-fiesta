package router

import (
	"github.com/labstack/echo"
	h "github.com/Gopher4U/sturdy-fiesta/RESTapi/handler"
)

func Routes(e *echo.Echo) {

	e.GET("/health", h.Health)
	e.POST("/luz/ligar", h.LigarLed)
	e.POST("/luz/desligar", h.DesligarLed)

}