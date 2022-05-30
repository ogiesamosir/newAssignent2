package main

import (
	"newassignmen2/db"
	"newassignmen2/routes"
)

func main() {
	db.StartDB()

	routes.StartRoute().Run(":8080")
}
