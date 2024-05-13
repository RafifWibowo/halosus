package main

import (
	"halosus/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.Connect()
}