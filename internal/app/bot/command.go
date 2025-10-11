package bot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func (b *Bot) handleCommand(ctx context.Context, text string, telegramUserID int64) string {
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
		dayNum, err := b.parseDayNumber(dayStr)
		if err != nil {
			return err.Error()
		}
		day := b.roadmap.GetByNumber(dayNum)
		if day == nil {
			log.Printf("⚠️  Failed get day %d in roadmap", dayNum)
			return "❌ Не удалось найти запись по roadmap. Попробуй позже."
		}
		return fmt.Sprintf("📅 День %d (Неделя %d)\n\n🔹 %s\n\n%s",
			day.Number, day.Week, day.Title, day.Description)

	case len(text) > 5 && strings.HasPrefix(text, "/done"):
		dayStr := strings.TrimSpace(text[5:])
		dayNum, err := b.parseDayNumber(dayStr)
		if err != nil {
			return err.Error()
		}
		user, err := b.userRepo.GetUserByTelegramID(ctx, telegramUserID)
		if err != nil {
			log.Printf("⚠️  Failed get user by telegram_id %d: %v", telegramUserID, err)
			return "❌ О тебе нет информации. Пройди регистрацию командой /start."
		}
		if err := b.progressRepo.MarkDayCompleted(ctx, user.ID, dayNum, b.roadmap.TotalDays); err != nil {
			log.Printf("⚠️  Failed mark day %d internal user_id %d: %v", dayNum, user.ID, err)
			return "❌ Не удалось сохранить прогресс. Попробуй позже."
		}
		return fmt.Sprintf("✅ День %d отмечен как выполненный!", dayNum)

	case text == "/progress":
		user, err := b.userRepo.GetUserByTelegramID(ctx, telegramUserID)
		if err != nil {
			log.Printf("⚠️  Failed get user by telegram_id %d: %v", telegramUserID, err)
			return "❌ О тебе нет информации. Пройди регистрацию командой /start."
		}
		completed, total, err := b.progressRepo.GetCompletionStats(ctx, user.ID, b.roadmap.TotalDays)
		if err != nil {
			log.Printf("⚠️  Failed get progress internal user_id %d: %v", user.ID, err)
			return "❌ Не удалось получить данные о прогрессе. Попробуй позже."
		}
		return fmt.Sprintf("📊 Прогресс: %d / %d дней выполнено.\n\nИспользуй /doneN, чтобы отмечать дни.", completed, total)

	default:
		return "❓❓❓\n\nЯ понимаю команды:\n/start — начать работу с ботом\n" +
			"/roadmap — персональный план изучения\n" +
			"/dayN — детали дня N\n/doneN - отметить выполнение дня N\n" +
			"/progress - полный прогресс"
	}
}

func (b *Bot) parseDayNumber(s string) (int, error) {
	if s == "" {
		return 0, errors.New("⚠️ Укажи номер дня, например: /done5, /day7")
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("⚠️ Номер дня должен быть целым числом")
	}
	if n < 1 || n > b.roadmap.TotalDays {
		return 0, fmt.Errorf("⚠️ Номер дня должен быть от 1 до %d", b.roadmap.TotalDays)
	}
	return n, nil
}
