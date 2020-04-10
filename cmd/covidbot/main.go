package main

import (
  "log";"github.com/go-telegram-bot-api/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Asia"),
    tgbotapi.NewKeyboardButton("Africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("America"),
    tgbotapi.NewKeyboardButton("Europe"),
  ),
)

func main() {
  bot, err := tgbotapi.NewBotAPI("1104178846:AAEG1Gl29ME2ONS-j9FPmMW10jzsaRa86to")
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

    switch update.Message.Text {
    case "open":
      msg.ReplyMarkup = numericKeyboard
    case "close":
      msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    }

    bot.Send(msg)
  }
}
