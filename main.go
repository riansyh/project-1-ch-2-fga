package main

import (
	"project-1/database"
	"project-1/routers"
)

func main() {
	var PORT = ":8081"

	database.StartDB()
	routers.StartServer().Run(PORT)
}
