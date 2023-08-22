package handler

import (
	"sync"
	"telegram-bot/internal/models"
)

func NewFilterPoll() models.FilterPoll {
	return models.FilterPoll{
		Mu:   sync.Mutex{},
		Poll: make(map[int64]models.Filter),
	}
}

func AddFilterForChat(p *models.FilterPoll, chatId int64) {
	p.Mu.Lock()
	filter := models.Filter{}
	p.Poll[chatId] = filter
	p.Mu.Unlock()
	return
}

func IsFilterExists(p *models.FilterPoll, chatId int64) bool {
	_, ok := p.Poll[chatId]
	if !ok {
		return false
	}
	return true
}

func DeleteFromPoll(p *models.FilterPoll, chatId int64) {
	p.Mu.Lock()
	delete(p.Poll, chatId)
	p.Mu.Unlock()
}
