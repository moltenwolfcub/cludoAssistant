package cludo

import (
	"testing"
)

func lookupOption(t *testing.T, category QuestionCategory, optionName string) (found *Option) {
	for _, o := range category.Options {
		if o.name == optionName {
			found = o
			return
		}
	}
	t.Errorf("Couldn't find option %s", optionName)
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
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteOption := lookupOption(t, game.whoCategory, "white")
	if !whiteOption.found {
		t.Error("Game.DoTurn() Was shown a who card but didn't update the who option")
	}

	pistolOption := lookupOption(t, game.whatCategory, "pistol")
	if pistolOption.found {
		t.Error("Game.DoTurn() Was shown a who card but the what option was set to found")
	}

	bedroomOption := lookupOption(t, game.whereCategory, "bedroom")
	if bedroomOption.found {
		t.Error("Game.DoTurn() Was shown a who card but the where option was set to found")
	}
}

func TestTurnWhatAnswerUpdatesFound(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolOption := lookupOption(t, game.whatCategory, "pistol")
	if !pistolOption.found {
		t.Error("Game.DoTurn() Was shown a what card but didn't update the what option")
	}

	whiteOption := lookupOption(t, game.whoCategory, "white")
	if whiteOption.found {
		t.Error("Game.DoTurn() Was shown a what card but the who option was set to found")
	}

	bedroomOption := lookupOption(t, game.whereCategory, "bedroom")
	if bedroomOption.found {
		t.Error("Game.DoTurn() Was shown a what card but the where option was set to found")
	}
}

func TestTurnWhereAnswerUpdatesFound(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomOption := lookupOption(t, game.whereCategory, "bedroom")
	if !bedroomOption.found {
		t.Error("Game.DoTurn() Was shown a where card but didn't update the where option")
	}

	whiteOption := lookupOption(t, game.whoCategory, "white")
	if whiteOption.found {
		t.Error("Game.DoTurn() Was shown a where card but the who option was set to found")
	}

	pistolOption := lookupOption(t, game.whatCategory, "pistol")
	if pistolOption.found {
		t.Error("Game.DoTurn() Was shown a where card but the what option was set to found")
	}
}

func TestTurnWhoAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhoAnswer)

	game.DoTurn(question)

	whiteOption := lookupOption(t, game.whoCategory, "white")
	if whiteOption.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a who card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhatAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhatAnswer)

	game.DoTurn(question)

	pistolOption := lookupOption(t, game.whatCategory, "pistol")
	if pistolOption.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a what card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestTurnWhereAnswerPosessor(t *testing.T) {
	game := GenSampleGame()

	question := NewQuestion(
		NewOption("white"),
		NewOption("pistol"),
		NewOption("bedroom"),
		"THIS",
		"alice",
	)
	question.SetAnswer(WhereAnswer)

	game.DoTurn(question)

	bedroomOption := lookupOption(t, game.whereCategory, "bedroom")
	if bedroomOption.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a where card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestStartingHandOneOption(t *testing.T) {
	game := GenSampleGame()
	game.AddStartingHand(
		[]*Option{NewOption("lead pipe")},
	)
	pipeOption := lookupOption(t, game.whatCategory, "lead pipe")

	if !pipeOption.found {
		t.Error("Game.AddStartingHand() Started with 1 card but it wasn't marked as found.")
	}
	if pipeOption.possessor != "THIS" {
		t.Error("Game.AddStartingHand() Started with 1 card but its owner wasn't THIS.")
	}
}
