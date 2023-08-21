package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"runtime/debug"
	"telegram-bot/internal/chain"
	"telegram-bot/internal/handler"
	"telegram-bot/internal/models"
	"telegram-bot/logging"
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

	//set chain
	taste := &chain.Taste{
		BaseAdviser: chain.BaseAdviser{
			Bot: bot,
		},
		H: handlers,
	}
	mood := &chain.Mood{
		BaseAdviser: chain.BaseAdviser{
			Bot: bot,
		},
	}
	mood.SetNext(taste)
	poll := &chain.Poll{
		BaseAdviser: chain.BaseAdviser{
			Bot: bot,
		},
	}
	poll.SetNext(mood)

	for update := range updates {
		//if update.Message != nil { // If we got a message
		//	handlers.Logger.Log.Info("Request to bot", zap.String("UserName", update.Message.From.UserName))
		//
		//	if update.Message == nil {
		//		continue
		//	}
		//
		//	if update.Message.IsCommand() {
		//		handler.HandleCommand(handlers, models.Bot, update.Message, auth, update)
		//	} else {
		//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlers.FileForUnknown.SendData())
		//		models.Bot.Send(msg)
		//	}
		//
		//}
		go processMsg(poll, update)
	}

	logger.Log.Info("Start bot...")
}
func processMsg(chain chain.MessageHandler, update tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err recover", err)
			fmt.Println("stacktrace from panic: ", string(debug.Stack()))
		}
	}()

	var chatID int64
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else if update.PollAnswer != nil {
		chatID = update.PollAnswer.User.ID
	} else {
		fmt.Println("failed")
		return
	}
	filter := &models.Filter{}
	chain.Execute(chatID, filter, update)
}

func HandleCommand(handler *handler.Handler, bot *tgbotapi.BotAPI, msg *tgbotapi.Message, ch chain.MessageHandler, update tgbotapi.Update) {
	//switch msg.Command() {
	//case "start":
	//	handler.Start(bot, msg.Chat.ID)
	//case "help":
	//	handler.Help(bot, msg.Chat.ID)
	//case "random":
	//	handler.RandomBeer(bot, msg.Chat.ID)
	//case "name":
	//	handler.BeerName(bot, msg.Chat.ID, []byte(msg.Text))
	//case "advice":
	//	go ch.Execute(msg.Chat.ID, &chain.Filter{}, update)
	//default:
	//	handler.UnknownReq(bot, msg.Chat.ID)
	//}
	ch.Execute(msg.Chat.ID, &models.Filter{}, update)
}
