package roadmap

type Day struct {
	Number      int
	Week        int
	Title       string
	Description string
}

func All() []Day {
	return []Day{
		{Number: 1, Week: 1, Title: "Установка Go и Hello World", Description: "Установи Go, настрой go mod, напиши 'Hello, [имя]!'"},
		{Number: 2, Week: 1, Title: "Переменные, типы, функции", Description: "Изучи базовые типы. Напиши программу для расчёта площади фигур."},
		{Number: 3, Week: 1, Title: "Управляющие конструкции", Description: "Реализуй калькулятор с проверкой деления на 0."},
		{Number: 4, Week: 1, Title: "Функции и множественный возврат", Description: "Напиши функции min/max и деления с остатком."},
		{Number: 5, Week: 1, Title: "Структуры и указатели", Description: "Создай структуру Person и функцию изменения возраста через указатель."},
		{Number: 6, Week: 1, Title: "Мини-проект: Бюджет", Description: "CLI-программа для учёта расходов (слайсы + структуры)."},
		{Number: 7, Week: 1, Title: "Отдых и чтение", Description: "Прочитай 'A Tour of Go: Basics'."},

		{Number: 8, Week: 2, Title: "Слайсы и мапы", Description: "Реализуй фильтрацию чисел и поиск столиц по странам."},
		{Number: 9, Week: 2, Title: "Работа с JSON", Description: "Сериализуй/десериализуй структуру User в файл."},
		{Number: 10, Week: 2, Title: "Обработка ошибок", Description: "Напиши функцию деления с возвратом ошибки."},
		{Number: 11, Week: 2, Title: "Тестирование (testing)", Description: "Напиши unit-тесты для своих функций."},
		{Number: 12, Week: 2, Title: "defer, panic, recover", Description: "Используй defer для закрытия файла, recover — для обработки паники."},
		{Number: 13, Week: 2, Title: "Мини-проект: CLI-утилита", Description: "Программа чтения JSON-файла с пользователями."},
		{Number: 14, Week: 2, Title: "Отдых и чтение", Description: "Прочитай 'Effective Go: Errors'."},

		{Number: 15, Week: 3, Title: "Goroutines", Description: "Запусти 3 параллельные функции."},
		{Number: 16, Week: 3, Title: "Каналы (channels)", Description: "Реализуй обмен данными между горутинами."},
		{Number: 17, Week: 3, Title: "sync.WaitGroup и Mutex", Description: "Защити счётчик от гонки."},
		{Number: 18, Week: 3, Title: "HTTP-сервер", Description: "Создай сервер с роутом /hello → JSON."},
		{Number: 19, Week: 3, Title: "HTTP-клиент и context", Description: "Сделай GET-запрос с таймаутом."},
		{Number: 20, Week: 3, Title: "Мини-проект: Счётчик запросов", Description: "HTTP-сервер с мьютексом и счётчиком."},
		{Number: 21, Week: 3, Title: "Отдых и чтение", Description: "Посмотри 'Go Concurrency Patterns' by Rob Pike."},

		{Number: 22, Week: 4, Title: "Интерфейсы", Description: "Создай интерфейс Notifier с Email/Sms реализациями."},
		{Number: 23, Week: 4, Title: "Встраивание (embedding)", Description: "Реализуй AdminUser через встраивание User."},
		{Number: 24, Week: 4, Title: "Dependency Injection", Description: "Переделай HTTP-сервер с DI."},
		{Number: 25, Week: 4, Title: "Линтеры и форматирование", Description: "Настрой golangci-lint."},
		{Number: 26, Week: 4, Title: "Generics", Description: "Напиши функцию Max[T] и стек Stack[T]."},
		{Number: 27, Week: 4, Title: "Мини-проект: Кэш", Description: "Реализуй интерфейс Cache с in-memory и mock."},
		{Number: 28, Week: 4, Title: "Отдых и чтение", Description: "Прочитай Go Proverbs."},

		{Number: 29, Week: 5, Title: "database/sql + SQLite", Description: "Создай таблицу users, вставь/выбери данные."},
		{Number: 30, Week: 5, Title: "ORM (GORM)", Description: "Перепиши работу с БД через GORM."},
		{Number: 31, Week: 5, Title: "Веб-фреймворк (Gin)", Description: "Создай REST API: GET/POST /users."},
		{Number: 32, Week: 5, Title: "Валидация и middleware", Description: "Добавь валидацию email и логирующий middleware."},
		{Number: 33, Week: 5, Title: "Тестирование HTTP", Description: "Напиши тесты через httptest."},
		{Number: 34, Week: 5, Title: "Мини-проект: REST API заметок", Description: "CRUD для заметок с сохранением в SQLite."},
		{Number: 35, Week: 5, Title: "Отдых и чтение", Description: "Изучи Gin examples на GitHub."},

		{Number: 36, Week: 6, Title: "Docker для Go", Description: "Напиши Dockerfile для API."},
		{Number: 37, Week: 6, Title: ".env и конфигурация", Description: "Используй godotenv для переменных."},
		{Number: 38, Week: 6, Title: "Миграции БД", Description: "Добавь миграцию через golang-migrate."},
		{Number: 39, Week: 6, Title: "Swagger", Description: "Добавь документацию API через swaggo."},
		{Number: 40, Week: 6, Title: "Финальная сборка", Description: "Проверь работоспособность всего."},
		{Number: 41, Week: 6, Title: "Финальный проект", Description: "Задеплой через Docker Compose."},
		{Number: 42, Week: 6, Title: "Рефлексия", Description: "Напиши итоги и планы на будущее."},
	}
}

func GetByNumber(dayNum int) *Day {
	for _, d := range All() {
		if d.Number == dayNum {
			return &d
		}
	}
	return nil
}
