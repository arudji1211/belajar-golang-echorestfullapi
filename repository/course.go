package repository

import (
	"context"

	"github.com/arudji1211/belajar-golang-echorestfullapi/core"
	"github.com/arudji1211/belajar-golang-echorestfullapi/usecase"
)

type Course struct {
	Id              string `json:"id"`
	NamaCourse      string `json:"nama_course"`
	DeskripsiCourse string `json:"deskripsi_course"`
	ThumbnailCourse string `json:"thumbnail_course"`
}

func (c *Course) Insert() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO course(nama_course,deskripsi_course,thumbnail_course) VALUES(?,?,?)", c.NamaCourse, c.DeskripsiCourse, c.ThumbnailCourse)
	if err != nil {
		panic(err)
	}

	return "Sukses Insert"
}

func (c *Course) Update() string {
	db := core.GetConnection()
	defer db.Close()

	mod := usecase.SqlLogic{}

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "UPDATE course SET nama_course=COALESCE(?,nama_course),deskripsi_course=COALESCE(?,deskripsi_course),thumbnail_course=COALESCE(?,thumbnail_course) WHERE id=?", mod.CreateNullString(c.NamaCourse), mod.CreateNullString(c.DeskripsiCourse), mod.CreateNullString(c.ThumbnailCourse), c.Id)
	if err != nil {
		panic(err)
	}

	return "Sukses Update"
}

func (c *Course) Delete() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "Delete FROM course WHERE id=?", c.Id)
	if err != nil {
		panic(err)
	}

	return "Sukses Delete"
}

func (c *Course) GetAll() []Course {
	var data []Course
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT * FROM course"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.NamaCourse, &c.DeskripsiCourse, &c.ThumbnailCourse)
		if err != nil {
			panic(err)
		}
		data = append(data, *c)
	}
	return data
}

func (c *Course) GetById() []Course {
	var data []Course
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id,nama_course,deskripsi_course,thumbnail_course FROM course WHERE id=?", c.Id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.NamaCourse, &c.DeskripsiCourse, &c.ThumbnailCourse)
		if err != nil {
			panic(err)
		}
		data = append(data, *c)
	}

	return data

}
