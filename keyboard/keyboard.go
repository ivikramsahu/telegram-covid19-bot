package keyboard

import(
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var AsiaKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Asia"),
    tgbotapi.NewKeyboardButton("Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("America"),
    tgbotapi.NewKeyboardButton("Europe"),
  ),
)

var AfricaKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("East Africa"),
    tgbotapi.NewKeyboardButton("West Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("North Africa"),
    tgbotapi.NewKeyboardButton("South Africa"),
  ),
)

var EuropeKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Asia"),
    tgbotapi.NewKeyboardButton("Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("America"),
    tgbotapi.NewKeyboardButton("Europe"),
  ),
)

var AmericaKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Asia"),
    tgbotapi.NewKeyboardButton("Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("America"),
    tgbotapi.NewKeyboardButton("Europe"),
  ),
)
