package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := application{port: os.Getenv("PORT")}

	app.serve()
}
