package models

import "github.com/arudji1211/belajar-golang-echorestfullapi/repository"

type Course struct {
}

func (c Course) GetAllCourse() []repository.Course {
	r := repository.Course{}
	return r.GetAll()
}

func (c Course) GetCourseById(id string) []repository.Course {
	r := repository.Course{}
	r.Id = id
	return r.GetById()
}

func (c Course) InsertCourse(nama_course, deskripsi_course, thumbnail_course string) string {
	r := repository.Course{}
	r.NamaCourse = nama_course
	r.DeskripsiCourse = deskripsi_course
	r.ThumbnailCourse = thumbnail_course
	return r.Insert()
}

func (c Course) UpdateCourse(id, nama_course, deskripsi_course, thumbnail_course string) string {
	r := repository.Course{}
	r.Id = id
	r.NamaCourse = nama_course
	r.DeskripsiCourse = deskripsi_course
	r.ThumbnailCourse = thumbnail_course
	return r.Update()
}

func (c Course) DeleteCourse(id string) string {
	r := repository.Course{}
	r.Id = id
	return r.Delete()
}
