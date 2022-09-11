package main

import (
	"github.com/arudji1211/belajar-golang-echorestfullapi/routes"
)

func main() {
	//start server
	e := routes.Router()
	e.Logger.Fatal(e.Start("localhost:1323"))
}
