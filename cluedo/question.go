package cluedo

type Answer int

const (
	UnknownAnswer Answer = iota
	WhoAnswer
	WhatAnswer
	WhereAnswer
	NoAnswer
)

type Question struct {
	whoPart   *Card
	whatPart  *Card
	wherePart *Card

	asker    *Player
	answerer *Player

	answer Answer
}

func NewQuestion(who, what, where *Card, asker, answerer *Player) Question {
	return Question{
		whoPart:   who,
		whatPart:  what,
		wherePart: where,
		asker:     asker,
		answerer:  answerer,
	}
}

func (q *Question) SetAnswer(a Answer) {
	q.answer = a
}

type Player struct {
	name      string
	cardCount int
}

func NewPlayer(name string, count int) *Player {
	return &Player{
		name:      name,
		cardCount: count,
	}
}
