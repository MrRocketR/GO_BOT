package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var token string = "7754890352:AAGAZ1tDIhMiwM59MUYwo9C1ZGk21cpfAoA"

// 👉 Задание 1: Создай структуру TeamScore!
// Подсказка: нам нужна мапа для хранения очков каждого участника
type TeamScore struct {
	scores map[string]int
}

// 👉 Задание 2: Напиши функцию создания нового счетчика
func NewTeamScore() *TeamScore {
	score := TeamScore{scores: make(map[string]int)}
	return &score
}

// 👉 Задание 3: Напиши функцию добавления очков
func (ts *TeamScore) AddScore(name string) {
	ts.scores[name]++
}

// 👉 Задание 4: Напиши функцию получения всех очков
func (ts *TeamScore) GetScores() string {
	var report string

	for k, v := range ts.scores {
		report += fmt.Sprintf("%s: %d \\n", k, v)

	}
	return report
}

func main() {
	// 👉 Задание 5: Инициализируй бота с твоим токеном
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("Ошибка при создании бота:", err)
	}

	// Создаем счетчик
	teamScore := NewTeamScore()

	// Настраиваем получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	fmt.Println("🤖 Бот запущен и ждет команд!")

	// 👉 Задание 6: Напиши обработку команд
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		switch text {
		case "/start":
			msg := tgbotapi.NewMessage(chatID, "Вас приветствует бот учета оценок в соревновании! Для просмотра оценок используйте /score. Для добавления нового участника и бала напишите /add [имя]")
			bot.Send(msg)
		case "/score":
			msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s: %s \\n", "Состав команды с очками", teamScore.GetScores()))
			bot.Send(msg)
		}

		// Твой код здесь!
		// Подсказки:
		// 1. Используй switch или if для проверки команд
		// 2. Команды: /start, /score, /add [имя]
		// 3. Для отправки сообщения используй:
		//    msg := tgbotapi.NewMessage(chatID, текст)
		//    bot.Send(msg)
	}
}
