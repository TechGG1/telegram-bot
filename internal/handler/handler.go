package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
	"strconv"
	"telegram-bot/internal/models"
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

func HandleCommand(handler *Handler, bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		handler.Start(bot, msg.Chat.ID)
	case "help":
		handler.Help(bot, msg.Chat.ID)
	case "random":
		handler.RandomBeer(bot, msg.Chat.ID)
	case "name":
		handler.BeerName(bot, msg.Chat.ID, []byte(msg.Text))
	default:
		handler.UnknownReq(bot, msg.Chat.ID)
	}
}

func (h *Handler) Ping(ctx *gin.Context) {
	chatid, err := strconv.ParseInt(ctx.Param("chatid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{
			"err": fmt.Sprint(err),
		})
		return
	}
	msgtext := fmt.Sprintf("AAAAAAAAAAAAAAAAAAAAA %d", chatid)
	msg := tgbotapi.NewMessage(chatid, msgtext)
	sendmsg, err := models.Bot.Send(msg)
	if err == nil {
		ctx.String(http.StatusOK, msgtext)
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     fmt.Sprint(err),
			"message": sendmsg,
		})
	}
}
