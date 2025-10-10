package bot

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/maximberdyshev/go-study-bot/internal/repository"
	"github.com/maximberdyshev/go-study-bot/internal/telegram"
)

const waitPollingTime = 2 * time.Second

type Bot struct {
	client        *telegram.Client
	userRepo      *repository.UserRepository
	progressRepo  *repository.ProgressRepository
	whitelistRepo *repository.WhitelistRepository
}

func New(
	client *telegram.Client,
	userRepo *repository.UserRepository,
	progressRepo *repository.ProgressRepository,
	whitelistRepo *repository.WhitelistRepository,
) *Bot {
	return &Bot{
		client:        client,
		userRepo:      userRepo,
		progressRepo:  progressRepo,
		whitelistRepo: whitelistRepo,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	log.Println("🤖  GoStudyBot launched in long polling mode")

	var offset int
	for {
		select {
		case <-ctx.Done():
			log.Println("🤖  GoStudyBot stopped")
			return ctx.Err()
		default:
		}

		updates, err := b.client.GetUpdates(ctx, offset)
		if err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("⚠️  Failed get updates: %v", err)
			time.Sleep(waitPollingTime)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			switch {
			case update.Message == nil:
				log.Printf("⚠️  Ignored without message")
				continue
			case update.Message.Text == "":
				log.Printf("⚠️  Ignored message without text")
				continue
			case update.Message.From == nil:
				log.Printf("⚠️  Ignored message from unknown user")
				continue
			}

			allowed, err := b.whitelistRepo.IsAllowed(ctx, update.Message.From.ID)
			if err != nil {
				log.Printf("⚠️  Failed to check whitelist for %d: %v", update.Message.From.ID, err)
				continue
			}
			if !allowed {
				log.Printf("🔒  Ignored unauthorized user: %d (@%s)",
					update.Message.From.ID,
					update.Message.From.Username)
				continue
			}

			user := &repository.User{
				TelegramID: int64(update.Message.From.ID),
				Username:   username(update.Message.From),
				FirstName:  update.Message.From.FirstName,
			}
			if err := b.userRepo.CreateUserIfNotExists(ctx, user); err != nil {
				log.Printf("⚠️  Failed save user %d: %v", user.TelegramID, err)
			}

			chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
			text := update.Message.Text

			log.Printf("📩  [%s] %s: %s", chatID, user.Username, text)

			response := b.handleCommand(ctx, text, user.TelegramID)
			if err := b.client.SendMessage(ctx, chatID, response); err != nil {
				log.Printf("❌  Failed send reply: %v", err)
			}
		}

		time.Sleep(waitPollingTime)
	}
}

func username(user *telegram.User) string {
	username := "user"
	if user.Username != "" {
		username = "@" + user.Username
	}
	return username
}
