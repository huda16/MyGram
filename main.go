package main

import (
	"mygram/database"
	"mygram/router"
	"os"
)

func main() {
	var PORT = os.Getenv("PORT")

	database.StartDB()
	router.StartServer().Run(PORT)
}
