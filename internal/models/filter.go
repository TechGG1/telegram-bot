package models

import "sync"

type Filter struct {
	Attr    []string
	IsMood  bool
	IsTaste bool
	IsSpec  bool
}

type FilterPoll struct {
	Mu   sync.Mutex
	Poll map[int64]Filter
}
