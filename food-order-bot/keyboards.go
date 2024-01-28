package main


import ( tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5" )


// MAIN MENU
var numeric_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Lavash", "menu:1"),
		tgbotapi.NewInlineKeyboardButtonData("Burger", "menu:2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Kebab", "menu:3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Pizza", "menu:4"),
		tgbotapi.NewInlineKeyboardButtonData("Drinks", "menu:5"),
	),
)


var lavash_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Lavash Mini", "menu:1:1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Lavash Classic", "menu:1:2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Lavash Big", "menu:1:3"),
	),
)


var burger_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Burger Classic", "menu:2:1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Burger Classic Cheese", "menu:2:2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Burger Big", "menu:2:3"),
	),
)


var kebab_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Kebab Beef", "menu:3:1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Kebab Chicken", "menu:3:2"),
	),
)


var pizza_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Pizza Small", "menu:4:1"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Pizza XXL", "menu:4:2"),
	),
)


var drinks_keyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Water 0.5", "menu:5:1"),
		tgbotapi.NewInlineKeyboardButtonData("Apple Juice Big", "menu:5:2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Coca-Cola 1.5", "menu:5:3"),
		tgbotapi.NewInlineKeyboardButtonData("Pepsi 1.5", "menu:5:4"),
		tgbotapi.NewInlineKeyboardButtonData("Sprite 1.5", "menu:5:5"),
	),
)