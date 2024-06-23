package cludo

import (
	"fmt"
	"slices"
)

type Option struct {
	name  string
	found bool

	possessor Player
}

func NewOption(name string) *Option {
	return &Option{
		name:  name,
		found: false,
	}
}

func (o *Option) SetFound() {
	o.found = true
}

type QuestionCategory struct {
	Options []*Option
}

func NewQuestionCategory(options ...*Option) QuestionCategory {
	return QuestionCategory{
		Options: options,
	}
}

func (q QuestionCategory) HasKnownSolution() bool {
	var available int
	for _, o := range q.Options {
		if !o.found {
			available++
		}
	}
	return available == 1
}

func (q *QuestionCategory) FoundOption(option *Option) {
	for _, o := range q.Options {
		if option.name == o.name {
			o.SetFound()
			return
		}
	}
}

type Answer int

const (
	WhoAnswer Answer = iota
	WhatAnswer
	WhereAnswer
	NoAnswer
)

type Question struct {
	whoPart   *Option
	whatPart  *Option
	wherePart *Option

	asker    Player
	answerer Player

	answer Answer
}

func NewQuestion(who, what, where *Option, asker, answerer Player) Question {
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

type Player string

type Game struct {
	whoCategory   QuestionCategory
	whatCategory  QuestionCategory
	whereCategory QuestionCategory

	players []Player
}

func NewDefaultGame(otherPlayers []string) Game {
	g := Game{
		whoCategory: NewQuestionCategory(
			NewOption("green"),
			NewOption("mustard"),
			NewOption("peacock"),
			NewOption("plum"),
			NewOption("scarlet"),
			NewOption("white"),
		),
		whatCategory: NewQuestionCategory(
			NewOption("wrench"),
			NewOption("candlestick"),
			NewOption("dagger"),
			NewOption("pistol"),
			NewOption("lead pipe"),
			NewOption("rope"),
		),
		whereCategory: NewQuestionCategory(
			NewOption("bathroom"),
			NewOption("study"),
			NewOption("dining room"),
			NewOption("games room"),
			NewOption("garage"),
			NewOption("bedroom"),
			NewOption("living room"),
			NewOption("kitchen"),
			NewOption("courtyard"),
		),
	}

	for _, other := range otherPlayers {
		p := Player(other)

		if p == "THIS" {
			panic("Can't have a player called `THIS`")
		}
		if slices.Contains(g.players, p) {
			panic("Can't have 2 players with the same name")
		}

		g.players = append(g.players, p)
	}
	g.players = append(g.players, "THIS")

	return g
}

func (g Game) EnsureValidQuestion(question Question) bool {
	if !slices.ContainsFunc(g.whoCategory.Options, func(o *Option) bool { return o.name == question.whoPart.name }) {
		return false
	}
	if !slices.ContainsFunc(g.whatCategory.Options, func(o *Option) bool { return o.name == question.whatPart.name }) {
		return false
	}
	if !slices.ContainsFunc(g.whereCategory.Options, func(o *Option) bool { return o.name == question.wherePart.name }) {
		return false
	}
	if !slices.Contains(g.players, question.asker) {
		return false
	}
	if !slices.Contains(g.players, question.answerer) {
		return false
	}

	return true
}

func (g *Game) DoTurn(question Question) {
	g.EnsureValidQuestion(question)

	switch question.answer {
	case NoAnswer:
		fmt.Println("Not Implemented")
	case WhoAnswer:
		g.whoCategory.FoundOption(question.whoPart)
	case WhatAnswer:
		g.whatCategory.FoundOption(question.whatPart)
	case WhereAnswer:
		g.whereCategory.FoundOption(question.wherePart)
	}
}
