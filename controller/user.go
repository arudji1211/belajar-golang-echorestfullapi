package controller

import (
	"net/http"

	"github.com/arudji1211/belajar-golang-echorestfullapi/models"

	"github.com/labstack/echo/v4"
)

type User struct {
}

var m models.User

func init() {
	m = models.User{}
}

func (a User) InsertData(c echo.Context) error {
	m.InsertUser(c.FormValue("username"), c.FormValue("password"), c.FormValue("nama_lengkap"))
	return c.String(http.StatusCreated, "OK")
}

func (a User) UpdateData(c echo.Context) error {
	m.UpdateUser(c.Param("id"), c.FormValue("nama_lengkap"), c.FormValue("password"), c.FormValue("username"))
	return c.String(http.StatusOK, "OK")
}

func (a User) DeleteData(c echo.Context) error {
	m.DeleteUser(c.Param("id"))
	return c.String(http.StatusOK, "OK")
}

func (a User) GetData(c echo.Context) error {
	data := m.GetUser(c.Param("id"))
	return c.JSON(http.StatusOK, data)
}

func (a User) GetAllData(c echo.Context) error {
	data := m.GetAllUser()
	return c.JSON(http.StatusOK, data)
}
