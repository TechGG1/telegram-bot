package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"telegram-bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (h *Handler) sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func (h *Handler) Start(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "Hi, I can advice you a delicious beer")
}

func (h *Handler) Help(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "Available commands:\n"+
		"/start - Start chatting with the bot\n"+
		"/help - Get a list of available commands\n"+
		"/random - Get a random delicious beer\n")
}

func (h *Handler) UnknownReq(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "ha-ha")
	msg := tgbotapi.NewMessage(chatID, h.FileForUnknown.SendData())
	bot.Send(msg)
}

func getBeer(r *http.Response) ([]models.Beer, error) {
	var arr []models.Beer
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &arr)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

func (h *Handler) RandomBeer(bot *tgbotapi.BotAPI, chatID int64) {
	resp, err := http.Get("https://api.punkapi.com/v2/beers/random")
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer: Curl random beer", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	beer, err := getBeer(resp)
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}

	beerBytes, err := json.Marshal(beer)
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer(Marshal)", zap.Error(err))
		return
	}
	msg := tgbotapi.NewMessage(chatID, string(beerBytes))
	bot.Send(msg)
}
