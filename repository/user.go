package repository

import (
	"context"

	"github.com/arudji1211/belajar-golang-echorestfullapi/core"
	"github.com/arudji1211/belajar-golang-echorestfullapi/usecase"
)

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NamaLengkap string `json:"nama_lengkap"`
}

func (c *User) Insert() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO user(username,password,nama_lengkap) VALUES(?,?,?)"
	_, err := db.ExecContext(ctx, script, c.Username, c.Password, c.NamaLengkap)
	if err != nil {
		panic(err)
	}

	return "Sukses Insert"

}

func (c *User) Update() string {
	db := core.GetConnection()
	defer db.Close()

	mod := usecase.SqlLogic{}
	ctx := context.Background()
	script := "UPDATE user set nama_lengkap=COALESCE(?,nama_lengkap),password=COALESCE(?,password),username=COALESCE(?,username) WHERE id=?"
	_, err := db.ExecContext(ctx, script, mod.CreateNullString(c.NamaLengkap), mod.CreateNullString(c.Password), mod.CreateNullString(c.Username), c.Id)
	if err != nil {
		panic(err)
	}
	return "Sukses Update"
}

func (c *User) Delete() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "DELETE from user WHERE id=?"
	_, err := db.ExecContext(ctx, script, c.Id)
	if err != nil {
		panic(err)
	}

	return "Sukses Delete"
}

func (c *User) GetById() []User {
	var data []User
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id,username,nama_lengkap FROM user WHERE id=?"
	rows, err := db.QueryContext(ctx, script, c.Id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Username, &c.NamaLengkap)
		if err != nil {
			panic(err)
		}
		data = append(data, *c)
	}
	return data
}

func (c *User) GetAll() []User {
	var data []User
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id,username,nama_lengkap,password FROM user"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Username, &c.NamaLengkap, &c.Password)
		if err != nil {
			panic(err)
		}
		data = append(data, *c)
	}
	return data
}
