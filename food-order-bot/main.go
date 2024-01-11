package main


import (
	"log"
	//"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var numeric_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)


func bot_start() {
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err != nil { log.Panic(err) }

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	log.Printf("authorized on account %s", bot.Self.UserName)

	for update := range updates {
		if update.Message == nil { continue }
		if !update.Message.IsCommand() { continue }


		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")


		switch update.Message.Command() {
			case "start":
				msg.Text = "This is a FOOD_ORDER_SYSTEM"

				msg.ReplyMarkup = numeric_keyboard

			case "menu":
				msg.Text = "[MENU]"
				//log.Printf("[MENU] %s", update.Message.Text)

			default: msg.Text = "[ERR] I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil { log.Panic(err) }
	}
}


func main() {
	bot_start()
}

