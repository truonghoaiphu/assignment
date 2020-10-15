package controllers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func (s *Server) initializeRoutes() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	// Home Route
	s.Router.HandleFunc("/", s.Home)
	s.Router.HandleFunc(os.Getenv("API_URL"), BroadcastMessage).Methods("POST")

}
