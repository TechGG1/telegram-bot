package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
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

	models.Bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		logger.Log.Error("Error in creating bot API", zap.Error(err))
	}

	models.Bot.Debug = true

	logger.Log.Info("Authorized on account", zap.String("bot info", models.Bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := models.Bot.GetUpdatesChan(u)

	handlers := handler.NewHandler(tgbotapi.FileURL(os.Getenv("UNKNOWN_COMMAND_MEM_URL")), logger)

	for update := range updates {
		if update.Message != nil { // If we got a message
			handlers.Logger.Log.Info("Request to bot", zap.String("UserName", update.Message.From.UserName))

			if update.Message == nil {
				continue
			}

			if update.Message.IsCommand() {
				handler.HandleCommand(handlers, models.Bot, update.Message)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlers.FileForUnknown.SendData())
				models.Bot.Send(msg)
			}

		}
	}
}

//func main() {
//	defaultLogLevel, err := zapcore.ParseLevel(os.Getenv("LOG_LEVEL"))
//	if err != nil {
//		log.Panic(err)
//	}
//	logger := logging.NewLogger(defaultLogLevel)
//
//	if err := godotenv.Load(); err != nil {
//		logger.Log.Error("Error in loading env", zap.Error(err))
//	}
//	token := os.Getenv("TELEGRAM_BOT_TOKEN")
//
//	models.Bot, err = tgbotapi.NewBotAPI(token)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	models.Bot.Debug = true
//
//	log.Printf("Authorized on account %s", models.Bot.Self.UserName)
//
//	wh, err := tgbotapi.NewWebhookWithCert("https://0.0.0.0:8443/"+models.Bot.Token, tgbotapi.FilePath("cert.pem"))
//	if err != nil {
//		panic(err)
//	}
//	_, err = models.Bot.Request(wh)
//	if err != nil {
//		panic(err)
//	}
//	info, err := models.Bot.GetWebhookInfo()
//	if err != nil {
//		panic(err)
//	}
//	if info.LastErrorDate != 0 {
//		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
//	}
//
//	updates := models.Bot.ListenForWebhook("/bot" + models.Bot.Token)
//	err = http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", http.HandlerFunc(HandleTelegramWebHook))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for update := range updates {
//		log.Printf("%+v\n", update)
//	}
//
//}
//
//func sayHi(ctx *gin.Context) {
//	log.Printf("Received GET")
//	chatid, err := strconv.ParseInt(ctx.Param("chatid"), 10, 64)
//	if err != nil {
//		ctx.JSON(http.StatusServiceUnavailable, gin.H{
//			"err": fmt.Sprint(err),
//		})
//		return
//	}
//	msgtext := fmt.Sprintf("AAAAAAAAAAAAAAAAAAAAA %d", chatid)
//	msg := tgbotapi.NewMessage(chatid, msgtext)
//	sendmsg, err := models.Bot.Send(msg)
//	if err == nil {
//		ctx.String(http.StatusOK, msgtext)
//	} else {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"err":     fmt.Sprint(err),
//			"message": sendmsg,
//		})
//	}
//}
//
//
//func parseTelegramRequest(r *http.Request) (*tgbotapi.Update, error) {
//	var update tgbotapi.Update
//	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
//		log.Printf("could not decode incoming update %s", err.Error())
//		return nil, err
//	}
//	return &update, nil
//}
//func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {
//
//	// Parse incoming request
//	var update, err = parseTelegramRequest(r)
//	if err != nil {
//		log.Printf("error parsing update, %s", err.Error())
//		return
//	}
//
//	// Send the punchline back to Telegram
//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "lol")
//	models.Bot.Send(msg)
//
//}
