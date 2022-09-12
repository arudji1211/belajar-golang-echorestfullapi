package controller

import (
	"net/http"

	"github.com/arudji1211/belajar-golang-echorestfullapi/models"
	"github.com/labstack/echo/v4"
)

type Course struct {
}

var mc models.Course

func init() {
	mc = models.Course{}
}

func (a Course) GetAllData(c echo.Context) error {
	data := mc.GetAllCourse()
	return c.JSON(http.StatusOK, data)
}

func (a Course) GetData(c echo.Context) error {
	data := m.GetUser(c.Param("id"))
	return c.JSON(http.StatusOK, data)
}

func (a Course) InsertData(c echo.Context) error {
	mc.InsertCourse(c.FormValue("nama_course"), c.FormValue("deskripsi_course"), c.FormValue("thumbnail_course"))
	return c.String(http.StatusCreated, "OK")
}

func (a Course) UpdateData(c echo.Context) error {
	mc.UpdateCourse(c.Param("id"), c.FormValue("nama_course"), c.FormValue("deskripsi_course"), c.FormValue("thumbnail_course"))
	return c.String(http.StatusOK, "OK")
}

func (a Course) DeleteData(c echo.Context) error {
	mc.DeleteCourse(c.Param("id"))
	return c.String(http.StatusOK, "OK")
}
