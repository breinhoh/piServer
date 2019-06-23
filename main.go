package main

import (
	"piServer/common/database"
	"piServer/routes"
)

func main() {

	database.Init()

	r := routes.Init()
	r.Run()
}
