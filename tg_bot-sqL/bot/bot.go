package bot

import (
	"log"
	"telegram-bot-sql-example/database"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
)

// Створюємо нового бота
func StartBot() error {
	bot, err := telego.NewBot("token")
	if err != nil {
		return err
	}


	// Отримуємо оновлення від користувачів(канал для повідомлень)
	updates, _ := bot.UpdatesViaLongPolling(nil)


	// Цикл обробки кожного повідомлення
	for update := range updates {
		if update.Message != nil {
			userID := update.Message.Chat.ID
			text := update.Message.Text

			// Якщо отримано команду /start, надсилаємо привітання
			if text == "/start" {
				msg := telegoutil.Message(telegoutil.ID(userID), "Вітаю! Введіть будь-яке повідомлення, і я його збережу.")
				bot.SendMessage(msg)
			} else if text == "/my_messages" {
				// Отримуємо всі повідомлення користувача
				messages, err := database.GetMessages(userID)
				if err != nil {
					log.Printf("Помилка отримання повідомлень: %v", err)
                    continue
				}

				// Формуємо відповідь з усіма повідомленнями
				response := "Ваші повідомлення:\n"
				for _, msg := range messages {
					response += "- " + msg + "\n"
				}

				msg := telegoutil.Message(telegoutil.ID(userID), response)
				bot.SendMessage(msg)
			} else {
				// Зберігаємо повідомлення в базі даних
				err := database.SaveMessage(userID, text)
				if err != nil {
					log.Printf("Помилка збереження повідомлення: %v", err)
				} else {
					msg := telegoutil.Message(telegoutil.ID(userID), "Ваше повідомлення збережено!")
					bot.SendMessage(msg)
				}
			}
		}
	}
	return nil
}
