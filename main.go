package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/tmsnn/goBot/telegram"
	"github.com/tmsnn/goBot/unsplash"
	"log"
	"os"
	"sync"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Loading .env error %s", err.Error())
	}

	ctx := context.TODO()
	wg := &sync.WaitGroup{}

	unsplashService := unsplash.NewService(os.Getenv("UNSPLASH_ACCESS_KEY"))
	telegramService := telegram.NewService(unsplashService)
	// loop over updates channel and handle incoming messages
	telegramService.GetUpdates(ctx, wg, os.Getenv("TELEGRAM_BOT_TOKEN"))
}