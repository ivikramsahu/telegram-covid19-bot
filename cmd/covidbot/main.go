package main

import (
  "log"
  dataredis "telegram-covid19-bot/dataRedis"
  "telegram-covid19-bot/keyboard"

  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
    case "/continents":
      msg.ReplyMarkup = continentKeyboard
    case "/aboutbot":
      msg.Text = keyboard.AboutBot
      msg.ParseMode = "HTML"
    case "/globalstatus":
      msg.Text = dataredis.GetDataSource(update.Message.Text + "@global")
      msg.ParseMode = "HTML"
    case "/asia":
      msg.Text = keyboard.AsiaKeyboard
      msg.ParseMode = "HTML"
    case "/africa":
      msg.Text = keyboard.AfricaKeyboard
      msg.ParseMode = "HTML"
    case "/america":
      msg.Text = keyboard.AmericaKeyboard
      msg.ParseMode = "HTML"
    case "/europe":
      msg.Text = keyboard.EuropeKeyboard
      msg.ParseMode = "HTML"
    case "/statewise":
      msg.Text = keyboard.StateKeyboard
      msg.ParseMode = "HTML"
    default:
      msg.Text = dataredis.GetDataSource(update.Message.Text)
      msg.ParseMode = "markdown"
    }
    if msg.Text == "" {
      msg.Text = "No data found"
    }
    bot.Send(msg)
  }
}
