package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/maximberdyshev/go-study-bot/internal/roadmap"
	"github.com/maximberdyshev/go-study-bot/internal/telegram"
)

const (
	waitPollingTime = 2 * time.Second
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN is required in .env")
	}

	client := telegram.NewClient(token)

	log.Println("🤖 GoStudyBot launched in long polling mode")

	if err := pollUpdates(context.Background(), client); err != nil {
		log.Fatalf("❌ Polling failed: %v", err)
	}
}

func pollUpdates(ctx context.Context, client *telegram.Client) error {
	var offset int

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		updates, err := client.GetUpdates(offset)
		if err != nil {
			log.Printf("⚠️ Failed to get updates: %v", err)
			time.Sleep(waitPollingTime)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil || update.Message.Text == "" {
				continue
			}

			// for debug
			adminID := os.Getenv("ADMIN_ID")
			if adminID == "" {
				log.Fatal("❌ ADMIN_ID is required in .env")
			} else {
				if update.Message.From == nil {
					log.Printf("🔒 Ignored message from unknown user")
					continue
				}
				allowID, err := strconv.ParseInt(adminID, 10, 64)
				if err != nil {
					log.Printf("❌ Invalid ADMIN_ID in .env: %v", err)
					continue
				}
				if update.Message.From.ID != int(allowID) {
					log.Printf("🔒 Ignored message from unallowed user: %d (@%s)",
						update.Message.From.ID,
						update.Message.From.Username)
					continue
				}
			}

			chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
			text := update.Message.Text
			username := "user"
			if update.Message.From != nil && update.Message.From.Username != "" {
				username = "@" + update.Message.From.Username
			}

			log.Printf("📩 [%s] %s: %s", chatID, username, text)

			responseText := handleCommand(text)
			if err := client.SendMessage(chatID, responseText); err != nil {
				log.Printf("❌ Failed to send reply: %v", err)
			}
		}

		time.Sleep(waitPollingTime)
	}
}

func handleCommand(text string) string {
	switch {
	case text == "/start":
		return "👋 Привет!\n Я GoStudyBot — твой персональный агент для изучения Go.\n\n" +
			"Напиши /roadmap, чтобы увидеть план или /dayN — чтобы узнать про день N."

	case text == "/roadmap":
		return "📚 Твой roadmap:\n\n" +
			"Неделя 1: Основы Go\n" +
			"Неделя 2: JSON, ошибки, тесты\n" +
			"Неделя 3: Конкурентность и HTTP\n" +
			"Неделя 4: Архитектура и интерфейсы\n" +
			"Неделя 5: БД и веб-API\n" +
			"Неделя 6: Docker и финальный проект\n\n" +
			"Напиши /day1, /day2, ... /day42"

	case len(text) > 4 && strings.HasPrefix(text, "/day"):
		dayStr := strings.TrimSpace(text[4:])
		if dayNum, err := strconv.Atoi(dayStr); err == nil {
			if day := roadmap.GetByNumber(dayNum); day != nil {
				return fmt.Sprintf("📅 День %d (Неделя %d)\n\n🔹 %s\n\n%s",
					day.Number, day.Week, day.Title, day.Description)
			}
			return "❌ День должен быть от 1 до 42."
		}
		return "❌ Неверный формат. Пример: /day7"

	default:
		return "❓ Я понимаю команды:\n/start — начать\n/roadmap — план\n/dayN — детали дня"
	}
}
