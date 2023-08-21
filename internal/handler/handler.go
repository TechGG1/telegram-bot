package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/logging"
)

type Handler struct {
	FileForUnknown tgbotapi.FileURL
	Logger         *logging.Logger
}

func NewHandler(memURL tgbotapi.FileURL, log *logging.Logger) *Handler {
	return &Handler{
		FileForUnknown: memURL,
		Logger:         log,
	}
}
