# 🤖 GoStudyBot

Telegram-бот, который помогает изучать Go по индивидуальному roadmap и отслеживать прогресс.

## 🚀 Функции
- ~~Персональный roadmap~~
- ~~Отметка выполненных дней~~
- ~~Хранение прогресса в PostgreSQL~~
- ~~Запуск через Docker Compose~~

## 🛠️ Как запустить (локально)

1. Создай бота через [@BotFather](https://t.me/BotFather) и получи токен.
2. Склонируй репозиторий:
   ```bash
   $> git clone https://github.com/maximberdyshev/go-study-bot.git .
   $> cd go-study-bot
   ```
3. Заполни .env файл.
4. Установи зависимости и запусти:
   ```bash
   $> go mod tidy
   $> go run cmd/bot/main.go
   ```
