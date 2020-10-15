package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"assignment/api/models"

	"github.com/joho/godotenv"
	"github.com/zhashkevych/scheduler"
)

func BroadcastMessage(w http.ResponseWriter, r *http.Request) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	t, _ := strconv.Atoi(os.Getenv("TIME"))
	ctx := context.Background()

	sc := scheduler.NewScheduler()
	sc.Add(ctx, sendMessage, time.Second*time.Duration(t))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}

func sendMessage(ctx context.Context) {
	messages := RandomMessage()
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	f, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(messages.Content)
}

func RandomMessage() models.Message {
	messages := models.Message{Content: models.RandomString(8), Time: time.Now()}
	return messages
}
