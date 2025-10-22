package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("8450193713:AAE_zTPT-Awxh_k_SMpp9dqYstmRj-VfyRw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	allowedChannels := map[int64]bool{
		-1002983106840: true,
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			if allowedChannels[chatID] {
				// Bot foydalanuvchi xabariga javob beradi
				msg := tgbotapi.NewMessage(chatID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				// Xabarni yuborish
				_, err := bot.Send(msg)
				if err != nil {
					log.Println("Send error:", err)
				}
				// Hech qachon o'chirmaydi
			}
		}
	}
}
