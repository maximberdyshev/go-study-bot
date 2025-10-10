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
3. Запусти Postgres локально:
   ```bash
   $> docker run --name go-study-db -e POSTGRES_PASSWORD=password -e POSTGRES_DB=study-bot -p 5432:5432 -d postgres:16
   ```
4. Заполни .env файл; корректно укажи secrets.
5. Установи зависимости и запусти:
   ```bash
   $> go mod tidy
   $> go run cmd/bot/main.go
   ```
