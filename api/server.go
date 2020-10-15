package api

import (
	"log"

	"assignment/api/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	server.Initialize()

	server.Run(":8080")

}
