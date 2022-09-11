package controller

import (
	"net/http"

	"github.com/arudji1211/belajar-golang-echorestfullapi/models"

	"github.com/labstack/echo/v4"
)

type Customer struct {
}

type response struct {
	Id           string
	Username     string
	Nama_lengkap string
}

var m models.User

func init() {
	m = models.User{}
}

func (a Customer) InsertData(c echo.Context) error {
	m.InsertUser(c.FormValue("username"), c.FormValue("password"), c.FormValue("nama_lengkap"))
	return c.String(http.StatusCreated, "Username:"+c.FormValue("username"))
}

func (a Customer) UpdateData(c echo.Context) error {
	resp := m.UpdateUser(c.Param("id"), c.FormValue("username"), c.FormValue("password"), c.FormValue("nama_lengkap"))
	return c.String(http.StatusOK, resp)
}

func (a Customer) DeleteData(c echo.Context) error {
	resp := m.DeleteUser(c.Param("id"))
	return c.String(http.StatusOK, resp)
}

func (a Customer) GetData(c echo.Context) error {
	data := m.GetUser(c.Param("id"))
	return c.JSON(http.StatusOK, data)
}

func (a Customer) GetAllData(c echo.Context) error {
	data := m.GetAllUser()
	return c.JSON(http.StatusOK, data)
}
