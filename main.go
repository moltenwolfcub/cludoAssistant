package main

import (
	"fmt"

	"github.com/moltenwolfcub/cluedoAssistant/cluedo"
)

func main() {
	game := cluedo.NewDefaultGame([]string{
		"alice",
		"bob",
		"charlie",
	})
	game.AddStartingHand(
		[]*cluedo.Card{
			cluedo.NewCard("peacock"),
			cluedo.NewCard("white"),
			cluedo.NewCard("rope"),
			cluedo.NewCard("bathroom"),
		},
	)

	q := cluedo.NewQuestion(
		cluedo.NewCard("scarlet"),
		cluedo.NewCard("dagger"),
		cluedo.NewCard("study"),
		"THIS",
		"alice",
	)
	q.SetAnswer(cluedo.WhoAnswer)
	game.DoTurn(q)

	q = cluedo.NewQuestion(
		cluedo.NewCard("green"),
		cluedo.NewCard("lead pipe"),
		cluedo.NewCard("games room"),
		"THIS",
		"alice",
	)
	q.SetAnswer(cluedo.NoAnswer)
	game.DoTurn(q)

	q = cluedo.NewQuestion(
		cluedo.NewCard("green"),
		cluedo.NewCard("lead pipe"),
		cluedo.NewCard("games room"),
		"THIS",
		"bob",
	)
	q.SetAnswer(cluedo.NoAnswer)
	game.DoTurn(q)

	q = cluedo.NewQuestion(
		cluedo.NewCard("green"),
		cluedo.NewCard("lead pipe"),
		cluedo.NewCard("games room"),
		"THIS",
		"charlie",
	)
	q.SetAnswer(cluedo.WhereAnswer)
	game.DoTurn(q)

	fmt.Println(game)
}
