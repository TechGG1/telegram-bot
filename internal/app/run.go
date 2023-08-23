package app

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
	"telegram-bot/internal/chain/filter"
	"telegram-bot/internal/handler"
	"telegram-bot/logging"
)

func Run() error {
	defaultLogLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Println(err)
		return err
	}
	logger := logging.NewLogger(defaultLogLevel)

	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error in loading env", zap.Error(err))
		return err
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		logger.Log.Error("Error in creating bot API", zap.Error(err))
		return err
	}

	bot.Debug = true

	logger.Log.Info("Authorized on account", zap.String("bot info", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	filters := filter.NewFilterPoll()

	handlers := handler.NewHandler(tgbotapi.FileURL(os.Getenv("UNKNOWN_COMMAND_MEM_URL")), logger, &filters)

	//set chain
	base := chain.BaseAdviser{
		Bot: bot,
	}
	taste := &chain.Taste{
		BaseAdviser: base,
		H:           handlers,
	}
	mood := &chain.Mood{
		BaseAdviser: base,
	}
	mood.SetNext(taste)
	poll := &chain.Poll{
		BaseAdviser: base,
	}
	poll.SetNext(mood)

	for update := range updates {
		if update.Message != nil || update.PollAnswer != nil {
			HandleCommand(handlers, bot, update.Message, poll, update)
		}
	}

	return nil
}

func HandleCommand(handler *handler.Handler, bot *tgbotapi.BotAPI, msg *tgbotapi.Message, ch chain.MessageHandler, update tgbotapi.Update) {
	if msg != nil && msg.IsCommand() {
		switch msg.Command() {
		case "start":
			handler.Start(bot, msg.Chat.ID)
		case "help":
			handler.Help(bot, msg.Chat.ID)
		case "random":
			handler.RandomBeer(bot, msg.Chat.ID)
		case "name":
			handler.BeerName(bot, msg.Chat.ID, msg.CommandArguments())
		case "advice":
			go processMsg(ch, update, handler)
		default:
			handler.UnknownReq(bot, msg.Chat.ID)
		}
	} else if update.PollAnswer != nil {
		go processMsg(ch, update, handler)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, handler.FileForUnknown.SendData())
		_, err := bot.Send(msg)
		if err != nil {
			handler.Logger.Log.Error("error in sendMessage", zap.Error(err))
		}
	}
}

func processMsg(ch chain.MessageHandler, update tgbotapi.Update, h *handler.Handler) {
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

	ch.Execute(chatID, h.Filter, update)
}
