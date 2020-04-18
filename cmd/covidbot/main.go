package main

import (
  "log"
  "telegram-covid19-bot/keyboard"
  "telegram-covid19-bot/dataRedis"
  "github.com/go-telegram-bot-api/telegram-bot-api"
)

var continentKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("/asia"),
    tgbotapi.NewKeyboardButton("/africa"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("/america"),
    tgbotapi.NewKeyboardButton("/europe"),
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
    case "/globalstatus":
      msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    case "/continents":
      msg.ReplyMarkup = continentKeyboard
    case "/coronavirus":
      msg.Text = keyboard.CoronaVirus
      msg.ParseMode = "markdown"
    case "/asia":
      msg.Text = keyboard.AsiaKeyboard
      msg.ParseMode = "markdown"
    case "/africa":
      msg.Text= keyboard.AfricaKeyboard
      msg.ParseMode = "markdown"
    case "/america":
      msg.Text = keyboard.AmericaKeyboard
      msg.ParseMode = "markdown"
    case "/europe":
      msg.Text = keyboard.EuropeKeyboard
      msg.ParseMode = "markdown"
    default:
      msg.Text = dataRedis.GetDataCountry(update.Message.Text)
      msg.ParseMode
    }

    bot.Send(msg)
  }
}
