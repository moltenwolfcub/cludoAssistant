package cluedo

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

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

func (q *QuestionCategory) FoundCard(card *Card, possessor *Player) (success bool) {
	for _, c := range q.Cards {
		if card.name == c.name {
			c.SetFound(possessor, true)
			return true
		}
	}
	return false
}

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

type Game struct {
	whoCategory   QuestionCategory
	whatCategory  QuestionCategory
	whereCategory QuestionCategory

	players []*Player
	Me      *Player
}

const MeIdent = "ME"

func NewDefaultGame(otherPlayers ...*Player) Game {
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

	g.Me = NewPlayer(MeIdent, 0)

	g.players = append(g.players, g.Me)
	for _, player := range otherPlayers {

		if player.name == MeIdent {
			panic(fmt.Sprintf("Can't have a player called `%s`", MeIdent))
		}
		if slices.Contains(g.players, player) {
			panic("Can't have 2 players with the same name")
		}

		g.players = append(g.players, player)
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
		if player.name == MeIdent {
			continue
		}
		playerList += fmt.Sprintf(" %v |", player.name)
		columnSpacing = append(columnSpacing, len(player.name))
	}

	allColumnWidth := len(playerList)

	// title
	left := int(math.Floor((float64(allColumnWidth) - 6) / 2))
	right := int(math.Ceil((float64(allColumnWidth) - 6) / 2))
	str := strings.Repeat("=", left) + " GAME " + strings.Repeat("=", right) + "\n"

	// player list
	str += playerList + "\n"

	sections := map[string][]*Card{
		"WHO ":   g.whoCategory.Cards,
		"WHAT ":  g.whatCategory.Cards,
		"WHERE ": g.whereCategory.Cards,
	}

	for title, cards := range sections {
		str += title + strings.Repeat("=", allColumnWidth-len(title)) + "\n"
		for _, card := range cards {
			str += fmt.Sprintf("| %s%s |", card.name, strings.Repeat(" ", columnSpacing[0]-len(card.name)))

			for i, width := range columnSpacing {
				if i == 0 {
					continue
				}

				str += " "

				if card.possessor == g.players[i-1] {
					str += "✓"
				} else if slices.Contains(card.nonPossessors, g.players[i-1]) {
					str += "x"
				} else {
					str += " "
				}

				str += strings.Repeat(" ", width)
				str += "|"
			}

			if card.IsFound() {
				str += " "
				if card.possessor.name != MeIdent {
					str += string(card.possessor.name)
				} else {
					str += "you"
				}
			}

			str += "\n"
		}

	}

	return str
}

func (g *Game) AddStartingHand(hand []*Card) {
	for _, c := range hand {
		if g.whoCategory.FoundCard(c, g.Me) {
			continue
		}
		if g.whatCategory.FoundCard(c, g.Me) {
			continue
		}
		if g.whereCategory.FoundCard(c, g.Me) {
			continue
		}
		panic(fmt.Sprintf("Can't have card `%v` in hand because it's not in the game.", c.name))
	}

	g.Me.cardCount = len(hand)

	for _, c := range g.whoCategory.Cards {
		if c.possessor != g.Me {
			c.nonPossessors = append(c.nonPossessors, g.Me)
		}
	}
	for _, c := range g.whatCategory.Cards {
		if c.possessor != g.Me {
			c.nonPossessors = append(c.nonPossessors, g.Me)
		}
	}
	for _, c := range g.whereCategory.Cards {
		if c.possessor != g.Me {
			c.nonPossessors = append(c.nonPossessors, g.Me)
		}
	}
	g.UpdateNonPossessors()
}

func (g *Game) UpdateNonPossessors() {
	allCards := []*Card{}
	allCards = append(allCards, g.whoCategory.Cards...)
	allCards = append(allCards, g.whatCategory.Cards...)
	allCards = append(allCards, g.whereCategory.Cards...)

	for _, c := range allCards {
		if !c.IsFound() {
			continue
		}
		for _, p := range g.players {
			if p == c.possessor {
				continue
			}
			c.AddNonPossessor(p)
		}
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
		g.analyseUnknownAnswer(question)
	case NoAnswer:
		for _, c := range g.whoCategory.Cards {
			if c.name == question.whoPart.name {
				c.AddNonPossessor(question.answerer)
				break
			}
		}
		for _, c := range g.whatCategory.Cards {
			if c.name == question.whatPart.name {
				c.AddNonPossessor(question.answerer)
				break
			}
		}
		for _, c := range g.whereCategory.Cards {
			if c.name == question.wherePart.name {
				c.AddNonPossessor(question.answerer)
				break
			}
		}
	case WhoAnswer:
		g.whoCategory.FoundCard(question.whoPart, question.answerer)
	case WhatAnswer:
		g.whatCategory.FoundCard(question.whatPart, question.answerer)
	case WhereAnswer:
		g.whereCategory.FoundCard(question.wherePart, question.answerer)
	}
	g.UpdateNonPossessors()
}

func (g *Game) analyseUnknownAnswer(question Question) {
	var gameWho, gameWhat, gameWhere *Card
	for _, c := range g.whoCategory.Cards {
		if c.name == question.whoPart.name {
			gameWho = c
			break
		}
	}
	for _, c := range g.whatCategory.Cards {
		if c.name == question.whatPart.name {
			gameWhat = c
			break
		}
	}
	for _, c := range g.whereCategory.Cards {
		if c.name == question.wherePart.name {
			gameWhere = c
			break
		}
	}

	// all the cards are already known so no new information can be gathered
	if gameWho.IsFound() && gameWhat.IsFound() && gameWhere.IsFound() {
		return
	}

	// if we already know the answerer has one of the cards in question we can't gather any useful information
	if gameWho.possessor == question.answerer || gameWhat.possessor == question.answerer || gameWhere.possessor == question.answerer {
		return
	}

	// i know the card isn't owned by the answerer
	whoFound := gameWho.IsFound() && gameWho.possessor != question.answerer
	whatFound := gameWhat.IsFound() && gameWhat.possessor != question.answerer
	whereFound := gameWhere.IsFound() && gameWhere.possessor != question.answerer

	// simple 2 knowns from question
	if whoFound && whatFound {
		gameWhere.SetFound(question.answerer, true)
		return
	}
	if whoFound && whereFound {
		gameWhat.SetFound(question.answerer, true)
		return
	}
	if whatFound && whereFound {
		gameWho.SetFound(question.answerer, true)
		return
	}

	// 1 known from question link creation
	if whoFound && !(gameWhere.IsFound() || gameWhat.IsFound()) {
		gameWhere.AddLink(question.answerer, gameWhat)
		gameWhat.AddLink(question.answerer, gameWhere)
		return
	}
	if whatFound && !(gameWhere.IsFound() || gameWho.IsFound()) {
		gameWhere.AddLink(question.answerer, gameWho)
		gameWho.AddLink(question.answerer, gameWhere)
		return
	}
	if whereFound && !(gameWho.IsFound() || gameWhat.IsFound()) {
		gameWho.AddLink(question.answerer, gameWhat)
		gameWhat.AddLink(question.answerer, gameWho)
		return
	}

	// no knowns from question
	if !(whoFound || whatFound || whereFound) {
		gameWho.AddTriLink(question.answerer, gameWhat, gameWhere)
		gameWhat.AddTriLink(question.answerer, gameWho, gameWhere)
		gameWhere.AddTriLink(question.answerer, gameWhat, gameWho)
	}
}
