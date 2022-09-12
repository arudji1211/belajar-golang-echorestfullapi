package routes

import (
	"github.com/arudji1211/belajar-golang-echorestfullapi/controller"
	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	//user
	cu := &controller.User{}
	e.POST("/user", cu.InsertData)
	e.PUT("/user/:id", cu.UpdateData)
	e.DELETE("/user/:id", cu.DeleteData)
	e.GET("/user/:id", cu.GetData)
	e.GET("/user", cu.GetAllData)

	cc := &controller.Course{}
	e.POST("/course", cc.InsertData)
	e.PUT("/course/:id", cc.UpdateData)
	e.DELETE("/course/:id", cc.DeleteData)
	e.GET("/course/:id", cc.GetData)
	e.GET("/course", cc.GetAllData)

	return e
}
