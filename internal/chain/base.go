package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Filter struct {
	Attr []string
}

type MessageHandler interface {
	Execute(int64, *Filter, tgbotapi.Update)
	SetNext(MessageHandler)
}

type BaseAdviser struct {
	Bot             *tgbotapi.BotAPI
	Next            MessageHandler
	isMoodSelected  bool
	isSpecSelected  bool
	isTasteSelected bool
}

func (b *BaseAdviser) SetNext(next MessageHandler) {
	b.Next = next
}

// SendMsg to telegram
func (b *BaseAdviser) SendMsg(c tgbotapi.Chattable) {
	if _, err := b.Bot.Send(c); err != nil {
		fmt.Println("err when send msg", err)
	}
}

// SomethingWentWrong send error to telegram
func (b *BaseAdviser) SomethingWentWrong(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Somening went wrong, try later")

	b.SendMsg(msg)
}
