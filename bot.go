package main

import (
	"log"
    "os"
    "fmt"
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

        var msg tgbotapi.MessageConfig
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
        if update.Message.NewChatMembers != nil{
            for index,user  := range *update.Message.NewChatMembers {
                log.Printf("[%d] users just joined the telegram group!",index)
                msg = welcome_user(update.Message, user)
                bot.Send(msg)
            }
        }	

        if update.Message.Text == "hola"{
            currentChat, err := bot.GetChat(update.Message.Chat.ChatConfig())
            log.Printf("[%s] %s", update.Message.From.UserName, err)
            fmt.Println(fmt.Sprintf(`El chat actual es de tipo --> %s ` ,currentChat.Type))
        }
    }
}
func welcome_user(msg *tgbotapi.Message , user tgbotapi.User) tgbotapi.MessageConfig {

    var username string = user.FirstName + " " + user.LastName
    new_msg := tgbotapi.NewMessage(msg.Chat.ID,fmt.Sprintf(` 🤖 Bienvenido al club de programación  %s 💻🤓 este grupo es para ayudar, guíar y proporcionar un espacio para discusión, estudio e innovación a todos los interesados en el mundo del software y/o ingeniería. Acercate a los mensajes fijados para leer las reglas 👌. 
    No olvides decir hola, compartir alguna duda, articulo o participar en algún proyecto que se esté llevando por miembros.`, username))
    return new_msg
}
