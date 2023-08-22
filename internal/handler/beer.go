package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"telegram-bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (h *Handler) sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		h.Logger.Log.Error("Error in sendMessage", zap.Error(err))
		return
	}
}

func (h *Handler) Start(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "Hi, I can advice you a delicious beer\n"+
		"Send /help to know about my features\n")
}

func (h *Handler) Help(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "Available commands:\n"+
		"/start - Start chatting with the bot\n"+
		"/help - Get a list of available commands\n"+
		"/random - Get a random delicious beer\n"+
		"/advice - ... \n")
}

func (h *Handler) UnknownReq(bot *tgbotapi.BotAPI, chatID int64) {
	h.sendMessage(bot, chatID, "ha-ha")
	msg := tgbotapi.NewMessage(chatID, h.FileForUnknown.SendData())
	_, err := bot.Send(msg)
	if err != nil {
		h.Logger.Log.Error("Error in sendMessage", zap.Error(err))
		return
	}
}

func getBeer(r *http.Response) ([]models.RespBeer, error) {
	var arr []models.RespBeer
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
	beerJson, err := json.MarshalIndent(beer, "", "  ")
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}
	strBeer := string(beerJson)
	strBeer = regexp.MustCompile(`[^a-zA-Z0-9:,.\/ \n]+`).ReplaceAllString(strBeer, "")
	h.sendMessage(bot, chatID, strBeer)
}

func (h *Handler) BeerName(bot *tgbotapi.BotAPI, chatID int64, name string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.punkapi.com/v2/beers/random", nil)
	if err != nil {
		h.Logger.Log.Error("error in FindBeerByParams: curl beer with params", zap.Error(err))
		return
	}
	q := req.URL.Query()
	q.Add("beer_name", name)
	randPage := rand.Intn(100)
	q.Add("page", strconv.Itoa(randPage))
	q.Add("per_page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		h.Logger.Log.Error("error in FindBeerByParams: curl beer with params", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	beer, err := getBeer(resp)
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}

	beerJson, err := json.MarshalIndent(beer, "", "  ")
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}
	strBeer := string(beerJson)
	strBeer = regexp.MustCompile(`[^a-zA-Z0-9:,.\/ \n]+`).ReplaceAllString(strBeer, "")
	h.sendMessage(bot, chatID, strBeer)
}

func (h *Handler) FindBeerByParams(bot *tgbotapi.BotAPI, chatID int64, params map[string]string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.punkapi.com/v2/beers/random", nil)
	if err != nil {
		h.Logger.Log.Error("error in FindBeerByParams: curl beer with params", zap.Error(err))
		return
	}

	q := req.URL.Query()
	q.Add("beer_name", params["beer_name"])
	q.Add("abv_gt", params["abv_gt"])
	randPage := rand.Intn(100)
	q.Add("page", strconv.Itoa(randPage))
	q.Add("per_page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		h.Logger.Log.Error("error in FindBeerByParams: curl beer with params", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	beer, err := getBeer(resp)
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}

	beerJson, err := json.MarshalIndent(beer, "", "  ")
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer", zap.Error(err))
		return
	}
	strBeer := string(beerJson)
	strBeer = regexp.MustCompile(`[^a-zA-Z0-9:,.\/ \n]+`).ReplaceAllString(strBeer, "")
	h.sendMessage(bot, chatID, strBeer)
}
