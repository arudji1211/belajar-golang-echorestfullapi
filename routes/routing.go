package routes

import (
	"github.com/arudji1211/belajar-golang-echorestfullapi/controller"
	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()
	c := &controller.User{}
	e.POST("/user", c.InsertData)
	e.PUT("/user/:id", c.UpdateData)
	e.DELETE("/user/:id", c.DeleteData)
	e.GET("/user/:id", c.GetData)
	e.GET("/user", c.GetAllData)

	return e
}
