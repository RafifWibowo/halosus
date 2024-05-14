package main

import (
	"halosus/db"
	"halosus/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.Connect()
	r := routes.Init()
	r.Run(":8080")
}