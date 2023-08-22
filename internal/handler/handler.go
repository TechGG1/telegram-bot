package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
	"telegram-bot/logging"
)

type Handler struct {
	FileForUnknown tgbotapi.FileURL
	Logger         *logging.Logger
	Filter         models.FilterPoll
}

func NewHandler(memURL tgbotapi.FileURL, log *logging.Logger, filter models.FilterPoll) *Handler {
	return &Handler{
		FileForUnknown: memURL,
		Logger:         log,
		Filter:         filter,
	}
}
