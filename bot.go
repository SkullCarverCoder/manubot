package main

import (
	"log"
    "os"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
        var msg tgbotapi.MessageConfig
        if update.Message.Text == "Hola"{
          msg = bienvenida_usuario(update.Message)
          bot.Send(msg)
        }
	

    }
}
func bienvenida_usuario(msg *tgbotapi.Message) tgbotapi.MessageConfig {

    var username string =msg.From.FirstName + " " + msg.From.LastName
    new_msg := tgbotapi.NewMessage(msg.Chat.ID,"Bienvenido al club de programación" + " " + username + "\n La aventura comienza")
    return new_msg
}
