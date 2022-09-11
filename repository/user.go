package repository

import (
	"context"

	"github.com/arudji1211/belajar-golang-echorestfullapi/core"
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

	script := "INSERT INTO user(username,password,nama_lengkap) VALUES('" + c.Username + "','" + c.Password + "','" + c.NamaLengkap + "')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	return "Sukses Insert"

}

func (c *User) Update() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "UPDATE user set username='" + c.Username + "',password='" + c.Password + "',nama_lengkap='" + c.Username + "' WHERE id='" + c.Id + "'"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	return "Sukses Update"
}

func (c *User) Delete() string {
	db := core.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "DELETE from user WHERE id='" + c.Id + "'"
	_, err := db.ExecContext(ctx, script)
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

	script := "SELECT id,username,nama_lengkap FROM user WHERE id='" + c.Id + "'"
	rows, err := db.QueryContext(ctx, script)
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

	script := "SELECT id,username,nama_lengkap FROM user"
	rows, err := db.QueryContext(ctx, script)
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
