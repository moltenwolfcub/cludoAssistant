package cludo

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Card struct {
	name string

	found     bool
	possessor Player
}

func NewCard(name string) *Card {
	return &Card{
		name:  name,
		found: false,
	}
}

func (c *Card) SetFound(possessor Player) {
	c.found = true
	c.possessor = possessor
}

type QuestionCategory struct {
	Cards []*Card
}

func NewQuestionCategory(cards ...*Card) QuestionCategory {
	return QuestionCategory{
		Cards: cards,
	}
}

func (q QuestionCategory) HasKnownSolution() bool {
	var available int
	for _, c := range q.Cards {
		if !c.found {
			available++
		}
	}
	return available == 1
}

func (q *QuestionCategory) FoundCard(card *Card, possessor Player) (success bool) {
	for _, c := range q.Cards {
		if card.name == c.name {
			c.SetFound(possessor)
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
	whoPart   *Card
	whatPart  *Card
	wherePart *Card

	asker    Player
	answerer Player

	answer Answer
}

func NewQuestion(who, what, where *Card, asker, answerer Player) Question {
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
			NewCard("green"),
			NewCard("mustard"),
			NewCard("peacock"),
			NewCard("plum"),
			NewCard("scarlet"),
			NewCard("white"),
		),
		whatCategory: NewQuestionCategory(
			NewCard("wrench"),
			NewCard("candlestick"),
			NewCard("dagger"),
			NewCard("pistol"),
			NewCard("lead pipe"),
			NewCard("rope"),
		),
		whereCategory: NewQuestionCategory(
			NewCard("bathroom"),
			NewCard("study"),
			NewCard("dining room"),
			NewCard("games room"),
			NewCard("garage"),
			NewCard("bedroom"),
			NewCard("living room"),
			NewCard("kitchen"),
			NewCard("courtyard"),
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

	longestCardNameLen := 0
	for _, c := range g.whoCategory.Cards {
		if l := len(c.name); l > longestCardNameLen {
			longestCardNameLen = l
		}
	}
	for _, c := range g.whatCategory.Cards {
		if l := len(c.name); l > longestCardNameLen {
			longestCardNameLen = l
		}
	}
	for _, c := range g.whereCategory.Cards {
		if l := len(c.name); l > longestCardNameLen {
			longestCardNameLen = l
		}
	}

	// player list setup
	playerList := "|  " + strings.Repeat(" ", longestCardNameLen) + "| you |"
	columnSpacing := []int{longestCardNameLen, 3}

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
	for _, card := range g.whoCategory.Cards {
		str += fmt.Sprintf("| %s%s |", card.name, strings.Repeat(" ", columnSpacing[0]-len(card.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if card.possessor == g.players[i-1] {
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
	for _, card := range g.whatCategory.Cards {
		str += fmt.Sprintf("| %s%s |", card.name, strings.Repeat(" ", columnSpacing[0]-len(card.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if card.possessor == g.players[i-1] {
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
	for _, card := range g.whereCategory.Cards {
		str += fmt.Sprintf("| %s%s |", card.name, strings.Repeat(" ", columnSpacing[0]-len(card.name)))

		for i, width := range columnSpacing {
			if i == 0 {
				continue
			}

			str += " "

			if card.possessor == g.players[i-1] {
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

func (g *Game) AddStartingHand(hand []*Card) {
	for _, c := range hand {
		if g.whoCategory.FoundCard(c, "THIS") {
			continue
		}
		if g.whatCategory.FoundCard(c, "THIS") {
			continue
		}
		if g.whereCategory.FoundCard(c, "THIS") {
			continue
		}
		panic(fmt.Sprintf("Can't have card `%v` in hand because it's not in the game.", c.name))
	}
}

func (g Game) EnsureValidQuestion(question Question) bool {
	if !slices.ContainsFunc(g.whoCategory.Cards, func(c *Card) bool { return c.name == question.whoPart.name }) {
		return false
	}
	if !slices.ContainsFunc(g.whatCategory.Cards, func(c *Card) bool { return c.name == question.whatPart.name }) {
		return false
	}
	if !slices.ContainsFunc(g.whereCategory.Cards, func(c *Card) bool { return c.name == question.wherePart.name }) {
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
		g.whoCategory.FoundCard(question.whoPart, question.answerer)
	case WhatAnswer:
		g.whatCategory.FoundCard(question.whatPart, question.answerer)
	case WhereAnswer:
		g.whereCategory.FoundCard(question.wherePart, question.answerer)
	}
}
