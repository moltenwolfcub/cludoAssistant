package cluedo

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	whoCategory   CardCategory
	whatCategory  CardCategory
	whereCategory CardCategory

	players []*Player
	Me      *Player
}

const MeIdent = "ME"

func NewDefaultGame(otherPlayers ...*Player) Game {
	g := Game{
		whoCategory: NewCardCategory(
			NewCard("green"),
			NewCard("mustard"),
			NewCard("peacock"),
			NewCard("plum"),
			NewCard("scarlet"),
			NewCard("white"),
		),
		whatCategory: NewCardCategory(
			NewCard("wrench"),
			NewCard("candlestick"),
			NewCard("dagger"),
			NewCard("pistol"),
			NewCard("lead pipe"),
			NewCard("rope"),
		),
		whereCategory: NewCardCategory(
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
	for _, c := range g.GetAllCards() {
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

	var str = strings.Builder{}
	str.WriteString(strings.Repeat("=", left) + " GAME " + strings.Repeat("=", right) + "\n")

	// player list
	str.WriteString(playerList + "\n")

	renderCategory := func(title string, cards []*Card) {
		str.WriteString(title + " " + strings.Repeat("=", allColumnWidth-len(title)-1) + "\n")
		for _, card := range cards {
			fmt.Fprintf(&str, "| %"+strconv.Itoa(columnSpacing[0])+"s |", card.name)

			for i, width := range columnSpacing {
				if i == 0 {
					continue
				}

				str.WriteString(" ")

				if card.possessor == g.players[i-1] {
					str.WriteString("✓")
				} else if slices.Contains(card.nonPossessors, g.players[i-1]) {
					str.WriteString("x")
				} else {
					str.WriteString(" ")
				}

				str.WriteString(strings.Repeat(" ", width))
				str.WriteString("|")
			}

			if card.IsFound() {
				str.WriteString(" ")
				if card.possessor.name != MeIdent {
					str.WriteString(string(card.possessor.name))
				} else {
					str.WriteString("you")
				}
			} else if card.isMurderItem {
				str.WriteString(" MURDER ELEMENT")
			}

			str.WriteString("\n")
		}
	}

	renderCategory("WHO", g.whoCategory.Cards)
	renderCategory("WHAT", g.whatCategory.Cards)
	renderCategory("WHERE", g.whereCategory.Cards)

	return str.String()
}

func (g Game) GetAllCards() []*Card {
	allCards := []*Card{}
	allCards = append(allCards, g.whoCategory.Cards...)
	allCards = append(allCards, g.whatCategory.Cards...)
	allCards = append(allCards, g.whereCategory.Cards...)
	return allCards
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

	for _, c := range g.GetAllCards() {
		if c.possessor != g.Me {
			c.nonPossessors = append(c.nonPossessors, g.Me)
		}
	}
	g.Update()
}

func (g *Game) Update() {
	g.UpdateCompleteCategories()
	g.UpdateNonPossessors()
	g.UpdateCompletePlayers()
}

func (g *Game) UpdateNonPossessors() {
	allCards := g.GetAllCards()

	for _, c := range allCards {
		if !c.IsFound() && !c.isMurderItem {
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

func (g *Game) UpdateCompletePlayers() {
	allCards := g.GetAllCards()

	for _, player := range g.players {
		if player == g.Me {
			continue
		}

		foundCardCount := 0
		unfoundCardCount := 0
		unknownCards := []*Card{}

		for _, c := range allCards {
			if slices.Contains(c.nonPossessors, player) {
				unfoundCardCount++
			} else if c.IsFound() {
				if c.possessor == player {
					foundCardCount++
				}
			}

			if !slices.Contains(c.nonPossessors, player) && c.possessor != player {
				unknownCards = append(unknownCards, c)
			}
		}

		if foundCardCount == player.cardCount {
			for _, c := range allCards {
				if c.possessor == player {
					continue
				}
				c.AddNonPossessor(player)
			}
		} else {
			cardsLeftToFind := player.cardCount - foundCardCount

			if len(unknownCards) == cardsLeftToFind {
				for _, c := range unknownCards {
					c.SetFound(player, true)
				}
			}
		}

	}
}

func (g *Game) UpdateCompleteCategories() {
	g.whoCategory.UpdateMurderKnowledge()
	g.whatCategory.UpdateMurderKnowledge()
	g.whereCategory.UpdateMurderKnowledge()

	for _, c := range g.GetAllCards() {
		if len(c.nonPossessors) == len(g.players) {
			c.isMurderItem = true
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

	// we already know our own cards so don't need to analyse
	if question.answerer == g.Me {
		return
	}

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
	g.Update()
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
