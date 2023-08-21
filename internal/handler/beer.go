package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	url "net/url"
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
		"/random - Get a random delicious beer\n"+
		"/name [args] - Get a list of available beer with specific name\n")
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
	h.sendMessage(bot, chatID, string(beerBytes))
}

func (h *Handler) BeerName(bot *tgbotapi.BotAPI, chatID int64, name []byte) {
	u, err := url.Parse("https://api.punkapi.com/v2")
	u.Scheme = "https"
	u.Host = "api.punkapi.com/v2"
	q := u.Query()
	q.Set("beer_name", string(name))
	q.Set("page", "1")
	q.Set("per_page", "1")
	u.RawQuery = q.Encode()
	fmt.Println("------------", u.RequestURI())
	resp, err := http.Get(u.RequestURI())
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

func (h *Handler) FindBeerByParams(bot *tgbotapi.BotAPI, chatID int64, params map[string]string) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.punkapi.com/v2/beers", nil)
	if err != nil {
		h.Logger.Log.Error("error in FindBeerByParams: curl beer with params", zap.Error(err))
		return
	}

	q := req.URL.Query()
	q.Add("beer_name", params["beer_name"])
	q.Add("abv_gt", params["abv_gt"])
	q.Add("page", "1")
	q.Add("per_page", "1")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	resp, err := client.Do(req)

	//resp, err := http.Get("https://api.punkapi.com/v2/beers/random")
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

	beerBytes, err := json.Marshal(beer)
	if err != nil {
		h.Logger.Log.Error("Error in RandomBeer(Marshal)", zap.Error(err))
		return
	}
	h.sendMessage(bot, chatID, string(beerBytes))
}
