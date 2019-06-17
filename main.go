package main

import (
	"piServer/common/database"
	"piServer/routes"
)

type TestDocument struct {
	Title string
	Note  string
}

func main() {

	database.Init()
	//db := database.GetDb()

	r := routes.Init()
	r.Run()
}
