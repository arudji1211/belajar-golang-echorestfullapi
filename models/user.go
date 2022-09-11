package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/arudji1211/belajar-golang-echorestfullapi/repository"
)

type User struct {
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u User) InsertUser(Username, Password, NamaLengkap string) string {
	r := repository.User{}
	r.Username = Username
	r.Password = GetMD5Hash(Password)
	r.NamaLengkap = NamaLengkap
	return r.Insert()
}

func (u User) UpdateUser(Id, Username, Password, NamaLengkap string) string {
	r := repository.User{}
	r.Id = Id
	r.Username = Username
	r.Password = GetMD5Hash(Password)
	r.NamaLengkap = NamaLengkap
	return r.Update()
}

func (u User) DeleteUser(Id string) string {
	r := repository.User{}
	r.Id = Id
	return r.Delete()
}

func (u User) GetAllUser() []repository.User {
	r := repository.User{}
	datas := r.GetAll()
	return datas
}

func (u User) GetUser(Id string) []repository.User {
	r := repository.User{}
	r.Id = Id
	datas := r.GetById()
	return datas
}
