package bot

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/maximberdyshev/go-study-bot/internal/repository"
	"github.com/maximberdyshev/go-study-bot/internal/roadmap"
	"github.com/maximberdyshev/go-study-bot/internal/telegram"
)

const (
	waitPollingTime = 2 * time.Second
)

type Config struct {
	AdminID string
}

type Bot struct {
	client *telegram.Client
	repo   *repository.UserRepository
	cfg    Config
}

func New(client *telegram.Client, repo *repository.UserRepository, cfg Config) *Bot {
	return &Bot{
		client: client,
		repo:   repo,
		cfg:    cfg,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	log.Println("🤖 GoStudyBot launched in long polling mode")

	var offset int
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		updates, err := b.client.GetUpdates(offset)
		if err != nil {
			log.Printf("⚠️ Failed get updates: %v", err)
			time.Sleep(waitPollingTime)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if update.Message == nil || update.Message.Text == "" {
				log.Printf("⚠️ Ignored message without text")
				continue
			}
			if update.Message.From == nil {
				log.Printf("⚠️ Ignored message from unknown user")
				continue
			}

			// debug
			if !b.isAdmin(update.Message.From) {
				log.Printf("🔒 Ignored message from unallowed user: %d (@%s)",
					update.Message.From.ID,
					update.Message.From.Username)
				continue
			}

			user := &repository.User{
				TelegramID: int64(update.Message.From.ID),
				Username:   username(update.Message.From),
				FirstName:  update.Message.From.FirstName,
			}
			if err := b.repo.CreateUserIfNotExists(ctx, user); err != nil {
				log.Printf("⚠️ Failed save user %d: %v", user.TelegramID, err)
			}

			chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
			text := update.Message.Text

			log.Printf("📩 [%s] %s: %s", chatID, user.Username, text)

			response := b.handleCommand(update.Message.Text)
			if err := b.client.SendMessage(chatID, response); err != nil {
				log.Printf("❌ Failed send reply: %v", err)
			}
		}

		time.Sleep(waitPollingTime)
	}
}

func (b *Bot) isAdmin(from *telegram.User) bool {
	adminID, err := strconv.ParseInt(b.cfg.AdminID, 10, 64)
	if err != nil {
		log.Printf("❌ Invalid ADMIN_ID in .env: %v", err)
		return false
	}
	return int64(from.ID) == adminID
}

func (b *Bot) handleCommand(text string) string {
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

func username(user *telegram.User) string {
	username := "user"
	if user.Username != "" {
		username = "@" + user.Username
	}
	return username
}
