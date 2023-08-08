package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"telegram-bot/internal/handler"
	"telegram-bot/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	defaultLogLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Panic(err)
	}
	logger := logging.NewLogger(defaultLogLevel)

	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error in loading env", zap.Error(err))
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		logger.Log.Error("Error in creating bot API", zap.Error(err))
	}

	bot.Debug = true

	logger.Log.Info("Authorized on account", zap.String("bot info", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	handlers := handler.NewHandler(tgbotapi.FileURL(os.Getenv("UNKNOWN_COMMAND_MEM_URL")), logger)

	for update := range updates {
		if update.Message != nil { // If we got a message
			handlers.Logger.Log.Info("Request to bot", zap.String("UserName", update.Message.From.UserName))

			if update.Message == nil {
				continue
			}

			if update.Message.IsCommand() {
				handler.HandleCommand(handlers, bot, update.Message)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlers.FileForUnknown.SendData())
				bot.Send(msg)
			}
		}
	}
}
