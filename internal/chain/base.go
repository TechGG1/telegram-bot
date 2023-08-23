package chain

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/internal/models"
)

type MessageHandler interface {
	Execute(int64, *models.FilterPoll, tgbotapi.Update)
	SetNext(MessageHandler)
}

type BaseAdviser struct {
	Bot  *tgbotapi.BotAPI
	Next MessageHandler
}

func (b *BaseAdviser) SetNext(next MessageHandler) {
	b.Next = next
}

func (b *BaseAdviser) sendPoll(chatID int64, question string, pollQuestions []string) error {
	// send poll
	poll := tgbotapi.NewPoll(chatID, question, pollQuestions...)
	poll.IsAnonymous = false
	poll.AllowsMultipleAnswers = false
	b.SendMsg(poll)

	return nil
}

func (b *BaseAdviser) SendMsg(c tgbotapi.Chattable) {
	if _, err := b.Bot.Send(c); err != nil {
		fmt.Println("err when send msg", err)
	}
}

func (b *BaseAdviser) SomethingWentWrong(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Somening went wrong, try later")

	b.SendMsg(msg)
}
