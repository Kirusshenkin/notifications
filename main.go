package main

import (
	"log"
	"os"

	 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Инициализация бота с токеном
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Настройка опций бота
	bot.Debug = true

	// Создание канала для получения обновлений от бота
	updates := bot.ListenForUpdates()

	// Обработка полученных обновлений
	go func() {
		for update := range updates {
			if update.Message == nil { // игнорируем обновления, не являющиеся сообщениями
				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Отвечаем на команду /release
			if update.Message.IsCommand() && update.Message.Command() == "release" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пора сделать релиз!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println("Ошибка отправки сообщения:", err)
				}
			}
		}
	}()

	log.Println("Бот запущен. Ожидание сообщений...")

	// Ожидание сигнала завершения
	select {}
}
