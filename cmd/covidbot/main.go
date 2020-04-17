package main

import (
  "fmt"
  "log"
  "telegram-covid19-bot/keyboard"
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
    fmt.Println(msg)
    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    switch update.Message.Text {
    case "/globalStatus":
      msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    case "/continents":
      msg.ReplyMarkup = continentKeyboard
    case "/asia":
      msg.ReplyToMessageID = keyboard.AsiaKeyboard
    case "/africa":
      msg.ReplyMarkup = keyboard.AfricaKeyboard
    case "/america":
      msg.ReplyMarkup = keyboard.AmericaKeyboard
    case "/europe":
      msg.ReplyMarkup = keyboard.EuropeKeyboard
    }

    bot.Send(msg)
  }
}
