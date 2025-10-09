package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/maximberdyshev/go-study-bot/internal/app/bot"
	"github.com/maximberdyshev/go-study-bot/internal/repository"
	"github.com/maximberdyshev/go-study-bot/internal/telegram"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	dbURL := os.Getenv("DATABASE_URL")
	if token == "" || dbURL == "" {
		log.Fatal("❌ Required env vars: TELEGRAM_BOT_TOKEN, DATABASE_URL")
	}

	db, err := repository.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("❌ Failed connect to DB: %v", err)
	}
	defer db.Close()
	log.Println("✅ Connected to DB")

	userRepo := repository.NewUserRepository(db)
	if err := userRepo.InitSchema(ctx); err != nil {
		log.Fatalf("❌ Failed initialization DB schema: %v", err)
	}
	log.Println("✅ Schema initialized")

	botApp := bot.New(
		telegram.NewClient(token),
		userRepo,
		bot.Config{AdminID: os.Getenv("ADMIN_ID")},
	)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("\n🔴 exit..")
		cancel()
	}()

	if err := botApp.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("❌ Bot failed: %v", err)
	}

	log.Println("🤖 GoStudyBot stopped")
}
