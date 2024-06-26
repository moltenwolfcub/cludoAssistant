package main

import (
	"fmt"

	"github.com/moltenwolfcub/cluedoAssistant/cluedo"
)

func main() {
	alice := cluedo.NewPlayer("alice", 0)
	bob := cluedo.NewPlayer("bob", 0)
	charlie := cluedo.NewPlayer("charlie", 0)

	game := cluedo.NewDefaultGame(
		alice,
		bob,
		charlie,
	)
	game.AddStartingHand(
		[]*cluedo.Card{
			cluedo.NewCard("peacock"),
			cluedo.NewCard("white"),
			cluedo.NewCard("rope"),
			cluedo.NewCard("bathroom"),
		},
	)

	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("white"),
			cluedo.NewCard("dagger"),
			cluedo.NewCard("study"),
			game.Me,
			alice,
		),
		cluedo.NoAnswer,
	)
	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("white"),
			cluedo.NewCard("dagger"),
			cluedo.NewCard("study"),
			game.Me,
			bob,
		),
		cluedo.NoAnswer,
	)
	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("white"),
			cluedo.NewCard("dagger"),
			cluedo.NewCard("study"),
			game.Me,
			charlie,
		),
		cluedo.WhatAnswer,
	)

	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("peacock"),
			cluedo.NewCard("lead pipe"),
			cluedo.NewCard("garage"),
			alice,
			bob,
		),
		cluedo.UnknownAnswer,
	)

	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("mustard"),
			cluedo.NewCard("lead pipe"),
			cluedo.NewCard("kitchen"),
			bob,
			charlie,
		),
		cluedo.NoAnswer,
	)
	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("mustard"),
			cluedo.NewCard("lead pipe"),
			cluedo.NewCard("kitchen"),
			bob,
			game.Me,
		),
		cluedo.NoAnswer,
	)
	AskQuestion(
		game,
		cluedo.NewQuestion(
			cluedo.NewCard("mustard"),
			cluedo.NewCard("lead pipe"),
			cluedo.NewCard("kitchen"),
			bob,
			alice,
		),
		cluedo.UnknownAnswer,
	)

	fmt.Println(game)
}

func AskQuestion(g cluedo.Game, q cluedo.Question, a cluedo.Answer) {
	q.SetAnswer(a)
	g.DoTurn(q)
}
