package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maximberdyshev/go-study-bot/internal/telegram"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN is required in .env")
	}

	adminID := os.Getenv("ADMIN_ID")
	if adminID == "" {
		log.Fatal("❌ ADMIN_ID is required in .env")
	}

	client := telegram.NewClient(token)
	err := client.SendMessage(adminID, "🚀 Привет!\n Это GoStudyBot — твой персональный агент для изучения Go!")
	if err != nil {
		log.Fatalf("❌ Failed to send message: %v", err)
	}

	fmt.Println("✅ Message has been sent")
}
