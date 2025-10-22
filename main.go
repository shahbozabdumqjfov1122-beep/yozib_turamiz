package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Bot token
	bot, err := tgbotapi.NewBotAPI("8450193713:AAE_zTPT-Awxh_k_SMpp9dqYstmRj-VfyRw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	// Ruxsat berilgan kanal
	allowedChannels := map[int64]bool{
		-1002983106840: true, // Shu yerga kanal ID
	}

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			// Faqat ruxsat berilgan kanallarga javob beradi
			if allowedChannels[chatID] {
				// Foydalanuvchi yozgan xabarni shunchaki reply qiladi
				msgText := update.Message.Text
				msg := tgbotapi.NewMessage(chatID, msgText)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
