package cludo

import (
	"testing"
)

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

	var whiteOption *Option
	for _, o := range game.whoCategory.Options {
		if o.name == "white" {
			whiteOption = o
			break
		}
	}
	if !whiteOption.found {
		t.Error("Game.DoTurn() Was shown a who card but didn't update the who option")
	}

	var pistolOption *Option
	for _, o := range game.whatCategory.Options {
		if o.name == "pistol" {
			pistolOption = o
			break
		}
	}
	if pistolOption.found {
		t.Error("Game.DoTurn() Was shown a who card but the what option was set to found")
	}

	var bedroomOption *Option
	for _, o := range game.whereCategory.Options {
		if o.name == "bedroom" {
			bedroomOption = o
			break
		}
	}
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

	var pistolOption *Option
	for _, o := range game.whatCategory.Options {
		if o.name == "pistol" {
			pistolOption = o
			break
		}
	}
	if !pistolOption.found {
		t.Error("Game.DoTurn() Was shown a what card but didn't update the what option")
	}

	var whiteOption *Option
	for _, o := range game.whoCategory.Options {
		if o.name == "white" {
			whiteOption = o
			break
		}
	}
	if whiteOption.found {
		t.Error("Game.DoTurn() Was shown a what card but the who option was set to found")
	}

	var bedroomOption *Option
	for _, o := range game.whereCategory.Options {
		if o.name == "bedroom" {
			bedroomOption = o
			break
		}
	}
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

	var bedroomOption *Option
	for _, o := range game.whereCategory.Options {
		if o.name == "bedroom" {
			bedroomOption = o
			break
		}
	}
	if !bedroomOption.found {
		t.Error("Game.DoTurn() Was shown a where card but didn't update the where option")
	}

	var whiteOption *Option
	for _, o := range game.whoCategory.Options {
		if o.name == "white" {
			whiteOption = o
			break
		}
	}
	if whiteOption.found {
		t.Error("Game.DoTurn() Was shown a where card but the who option was set to found")
	}

	var pistolOption *Option
	for _, o := range game.whatCategory.Options {
		if o.name == "pistol" {
			pistolOption = o
			break
		}
	}
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

	var whiteOption *Option
	for _, o := range game.whoCategory.Options {
		if o.name == "white" {
			whiteOption = o
			break
		}
	}
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

	var pistolOption *Option
	for _, o := range game.whatCategory.Options {
		if o.name == "pistol" {
			pistolOption = o
			break
		}
	}
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

	var bedroomOption *Option
	for _, o := range game.whereCategory.Options {
		if o.name == "bedroom" {
			bedroomOption = o
			break
		}
	}
	if bedroomOption.possessor != "alice" {
		t.Error("Game.DoTurn() Was shown a where card by alice but she wasn't marked as the owner of the card.")
	}
}

func TestStartingHandOneOption(t *testing.T) {
	game := GenSampleGame()
	game.AddStartingHand(
		[]*Option{NewOption("lead pipe")},
	)
	var pipeOption *Option
	for _, o := range game.whatCategory.Options {
		if o.name == "lead pipe" {
			pipeOption = o
			break
		}
	}

	if !pipeOption.found {
		t.Error("Game.AddStartingHand() Started with 1 card but it wasn't marked as found.")
	}
	if pipeOption.possessor != "THIS" {
		t.Error("Game.AddStartingHand() Started with 1 card but its owner wasn't THIS.")
	}
}
