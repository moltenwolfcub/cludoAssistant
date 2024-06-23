package cludo

import (
	"slices"
	"testing"
)

func lookupCard(t *testing.T, category QuestionCategory, cardName string) (found *Card) {
	for _, c := range category.Cards {
		if c.name == cardName {
			found = c
			return
		}
	}
	t.Errorf("Couldn't find card %s", cardName)
	return nil
}

func GenSampleGame() Game {
	players := []string{
		"alice",
		"bob",
		"charlie",
	}
	return NewDefaultGame(players)
}

func TestTurnWhoAnswerUpdatesFound(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if !whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but didn't update the who card")
	}

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but the what card was set to found")
	}

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a who card but the where card was set to found")
	}
}

func TestTurnWhatAnswerUpdatesFound(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if !pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but didn't update the what card")
	}

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but the who card was set to found")
	}

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a what card but the where card was set to found")
	}
}

func TestTurnWhereAnswerUpdatesFound(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if !bedroomCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but didn't update the where card")
	}

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but the who card was set to found")
	}

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.IsFound() {
		t.Error("Game.DoTurn() Was shown a where card but the what card was set to found")
	}
}

func TestTurnWhoAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteCard := lookupCard(t, game.whoCategory, "white")
	if whiteCard.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a who card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhatAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolCard := lookupCard(t, game.whatCategory, "pistol")
	if pistolCard.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a what card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhereAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomCard := lookupCard(t, game.whereCategory, "bedroom")
	if bedroomCard.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a where card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestStartingHandOneCard(t *testing.T) {
	game := GenSampleGame()
	game.AddStartingHand(
		[]*Card{NewCard("lead pipe")},
	)

	pipeCard := lookupCard(t, game.whatCategory, "lead pipe")
	if !pipeCard.IsFound() {
		t.Error("Game.AddStartingHand() Started with 1 card but it wasn't marked as found.")
	}
	if pipeCard.possessor != "THIS" {
		t.Error("Game.AddStartingHand() Started with 1 card but its owner wasn't THIS.")
	}
}

func TestStartingHandMultipleCards(t *testing.T) {
	game := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("wrench"),
		NewCard("green"),
		NewCard("study"),
		NewCard("bathroom"),
	})

	if !lookupCard(t, game.whatCategory, "wrench").IsFound() {
		t.Error("Game.AddStartingHand() Wrench card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whoCategory, "green").IsFound() {
		t.Error("Game.AddStartingHand() Green card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whereCategory, "study").IsFound() {
		t.Error("Game.AddStartingHand() Study card wasn't marked as found when it was in the starting hand")
	}
	if !lookupCard(t, game.whereCategory, "bathroom").IsFound() {
		t.Error("Game.AddStartingHand() Bathroom card wasn't marked as found when it was in the starting hand")
	}
}

func TestTurnNonPossessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewCard("white"),
		NewCard("pistol"),
		NewCard("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(NoAnswer)

	game.DoTurn(question)

	whoCard := lookupCard(t, game.whoCategory, "white")
	if !slices.Contains(whoCard.nonPossessors, "alice") {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the person")
	}
	whatCard := lookupCard(t, game.whatCategory, "pistol")
	if !slices.Contains(whatCard.nonPossessors, "alice") {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the weapon")
	}
	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if !slices.Contains(whereCard.nonPossessors, "alice") {
		t.Error("Game.DoTurn() Alice couldn't answer the question but she wasn't marked as not having the location")
	}
}

func TestUnkownAnswerWith2SimpleKnown(t *testing.T) {
	game := GenSampleGame()
	game.AddStartingHand([]*Card{
		NewCard("green"),
		NewCard("dagger"),
	})

	question := NewQuestion(
		NewCard("green"),
		NewCard("dagger"),
		NewCard("bedroom"),
		"bob",
		"alice",
	)
	question.SetAnswer(UnknownAnswer)

	game.DoTurn(question)

	whereCard := lookupCard(t, game.whereCategory, "bedroom")
	if whereCard.possessor != "alice" {
		t.Error("Game.analyseUnknownAnswer() I had 2 cards and alice showed a card when asked about them but the 3rd wasn't marked as hers.")
	}
}
