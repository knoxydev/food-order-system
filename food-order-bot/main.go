package main


import (
	"log"
	//"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
				case "start":
					msg.Text = "This is a FOOD_ORDER_SYSTEM"
					msg.ReplyMarkup = numeric_keyboard

				case "menu": msg.Text = "[MENU]"
				default: msg.Text = "[ERR] I don't know that command"
			}

			if _, err := bot.Send(msg); err != nil { log.Panic(err) }

		} else if update.CallbackQuery != nil {

			// Respond to the callback query, telling Telegram to show the user a message with the data received
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil { panic(err) }
			// And finally, send a message containing the data received.
			//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			//if _, err := bot.Send(msg); err != nil { panic(err) }


			some_kyb := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

			switch update.CallbackQuery.Data {
				case "menu:1":
					some_kyb.Text = "➜ LAVASH"
					some_kyb.ReplyMarkup = lavash_keyboard
				case "menu:2":
					some_kyb.Text = "➜ BURGERS"
					some_kyb.ReplyMarkup = burger_keyboard
				case "menu:3":
					some_kyb.Text = "➜ KEBAB"
					some_kyb.ReplyMarkup = kebab_keyboard
				case "menu:4":
					some_kyb.Text = "➜ PIZZA"
					some_kyb.ReplyMarkup = pizza_keyboard
				case "menu:5":
					some_kyb.Text = "➜ DRINKS"
					some_kyb.ReplyMarkup = drinks_keyboard
			}

			if _, err := bot.Send(some_kyb); err != nil { log.Panic(err) }

			// log.Printf("\n--------------")
			// log.Println(update.CallbackQuery.Data)
			// log.Printf("--------------\n")
		}

	}
}


func main() {
	bot_start()
}

