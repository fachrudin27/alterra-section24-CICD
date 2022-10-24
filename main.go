package main

import (
	"log"
	"os"
	"praktikum/databases"
	"praktikum/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databases.InitDB()

	e := echo.New()

	router.New(e, databases.DB)

	port := os.Getenv("START")

	e.Start(port)
}
