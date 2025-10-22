package main

import (
	"log"
	"time"

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
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID

			if allowedChannels[chatID] {
				// Bot foydalanuvchi xabarini reply qiladi
				msg := tgbotapi.NewMessage(chatID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				sentMsg, err := bot.Send(msg)
				if err != nil {
					log.Println("Send error:", err)
					continue
				}

				// 5 soniya keyin bot o'z xabarini o'chiradi
				go func(chatID int64, messageID int) {
					time.Sleep(5 * time.Second)
					delMsg := tgbotapi.DeleteMessageConfig{
						ChatID:    chatID,
						MessageID: messageID,
					}
					if _, err := bot.Request(delMsg); err != nil {
						log.Println("Delete error:", err)
					}
				}(chatID, sentMsg.MessageID)
			}
		}
	}
}
