package cludo

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Option struct {
	name string

	found     bool
	possessor Player
}

func NewOption(name string) *Option {
	return &Option{
		name:  name,
		found: false,
	}
}

func (o *Option) SetFound(possessor Player) {
	o.found = true
	o.possessor = possessor
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

func (q *QuestionCategory) FoundOption(option *Option, possessor Player) (success bool) {
	for _, o := range q.Options {
		if option.name == o.name {
			o.SetFound(possessor)
			return true
		}
	}
	return false
}

type Answer int

const (
	WhoAnswer Answer = iota
	WhatAnswer
	WhereAnswer
	UnknownAnswer
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

	g.players = append(g.players, "THIS")
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

	return g
}

func (g Game) String() string {
	// setup

	longestOptionLen := 0
	for _, o := range g.whoCategory.Options {
		if l := len(o.name); l > longestOptionLen {
			longestOptionLen = l
		}
	}
	for _, o := range g.whatCategory.Options {
		if l := len(o.name); l > longestOptionLen {
			longestOptionLen = l
		}
	}
	for _, o := range g.whereCategory.Options {
		if l := len(o.name); l > longestOptionLen {
			longestOptionLen = l
		}
	}

	// player list setup
	playerList := "|  " + strings.Repeat(" ", longestOptionLen) + "| you |"
	columnSpacing := []int{longestOptionLen, 3}

	for _, player := range g.players {
		if player == "THIS" {
			continue
		}
		playerList += fmt.Sprintf(" %v |", player)
		columnSpacing = append(columnSpacing, len(player))
	}

	allColumnWidth := len(playerList)

	// title
	left := int(math.Floor((float64(allColumnWidth) - 6) / 2))
	right := int(math.Ceil((float64(allColumnWidth) - 6) / 2))
	str := strings.Repeat("=", left) + " GAME " + strings.Repeat("=", right) + "\n"

	// player list
	str += playerList + "\n"

	// WHO
	str += "WHO " + strings.Repeat("=", allColumnWidth-4) + "\n"
	for _, option := range g.whoCategory.Options {
		str += fmt.Sprintf("| %s%s |", option.name, strings.Repeat(" ", columnSpacing[0]-len(option.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if option.possessor == g.players[i-1] {
				str += "✓"
			} else {
				str += " "
			}

			str += strings.Repeat(" ", width)
			str += "|"
		}

		str += "\n"
	}

	// WHAT
	str += "WHAT " + strings.Repeat("=", allColumnWidth-5) + "\n"
	for _, option := range g.whatCategory.Options {
		str += fmt.Sprintf("| %s%s |", option.name, strings.Repeat(" ", columnSpacing[0]-len(option.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if option.possessor == g.players[i-1] {
				str += "✓"
			} else {
				str += " "
			}

			str += strings.Repeat(" ", width)
			str += "|"
		}

		str += "\n"
	}

	// WHERE
	str += "WHERE " + strings.Repeat("=", allColumnWidth-6) + "\n"
	for _, option := range g.whereCategory.Options {
		str += fmt.Sprintf("| %s%s |", option.name, strings.Repeat(" ", columnSpacing[0]-len(option.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if option.possessor == g.players[i-1] {
				str += "✓"
			} else {
				str += " "
			}

			str += strings.Repeat(" ", width)
			str += "|"
		}

		str += "\n"
	}

	return str
}

func (g *Game) AddStartingHand(hand []*Option) {
	for _, o := range hand {
		if g.whoCategory.FoundOption(o, "THIS") {
			continue
		}
		if g.whatCategory.FoundOption(o, "THIS") {
			continue
		}
		if g.whereCategory.FoundOption(o, "THIS") {
			continue
		}
		panic(fmt.Sprintf("Can't have card `%v` in hand because it's not in the game.", o.name))
	}
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
	case UnknownAnswer:
		fmt.Println("Not Implemented")
	case NoAnswer:
		fmt.Println("Not Implemented")
	case WhoAnswer:
		g.whoCategory.FoundOption(question.whoPart, question.answerer)
	case WhatAnswer:
		g.whatCategory.FoundOption(question.whatPart, question.answerer)
	case WhereAnswer:
		g.whereCategory.FoundOption(question.wherePart, question.answerer)
	}
}
