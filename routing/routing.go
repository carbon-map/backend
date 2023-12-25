package routing

import (
	"backend/controller" //引入controller套件

	"github.com/labstack/echo/v4" //引入echo框架
)

func Routing(e *echo.Echo) {
	e.GET("/", controller.Home)
	e.GET("/data", controller.DataServe)
	e.GET("/dataInterval", controller.DataRange)
}
