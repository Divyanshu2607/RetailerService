package main

import (
	"RetailerAPI/database"
	"RetailerAPI/routes"
)

func main() {
	database.InitDB() //initialising data base

	r := routes.SetUpRouter()

	r.Run(":8080") //starting server
}
