package controller

import (
	"github.com/labstack/echo/v4"
)

func Routings() {
	e := echo.New()
	c := &Customer{}
	e.POST("/user", c.InsertData)
	e.PUT("/user/:id", c.UpdateData)
	e.DELETE("/user/:id", c.DeleteData)
	e.GET("/user/:id", c.GetData)
	e.GET("/user", c.GetAllData)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
