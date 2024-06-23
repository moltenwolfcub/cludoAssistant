package main

import (
	"fmt"

	"github.com/moltenwolfcub/cludoAssistant/cludo"
)

func main() {
	game := cludo.NewDefaultGame([]string{
		"alice",
		"bob",
		"charlie",
	})
	game.AddStartingHand(
		[]*cludo.Card{
			cludo.NewCard("peacock"),
			cludo.NewCard("white"),
			cludo.NewCard("rope"),
			cludo.NewCard("bathroom"),
		},
	)

	q := cludo.NewQuestion(
		cludo.NewCard("scarlet"),
		cludo.NewCard("dagger"),
		cludo.NewCard("study"),
		"THIS",
		"alice",
	)
	q.SetAnswer(cludo.WhoAnswer)
	game.DoTurn(q)

	q = cludo.NewQuestion(
		cludo.NewCard("green"),
		cludo.NewCard("lead pipe"),
		cludo.NewCard("games room"),
		"THIS",
		"alice",
	)
	q.SetAnswer(cludo.NoAnswer)
	game.DoTurn(q)

	q = cludo.NewQuestion(
		cludo.NewCard("green"),
		cludo.NewCard("lead pipe"),
		cludo.NewCard("games room"),
		"THIS",
		"bob",
	)
	q.SetAnswer(cludo.NoAnswer)
	game.DoTurn(q)

	q = cludo.NewQuestion(
		cludo.NewCard("green"),
		cludo.NewCard("lead pipe"),
		cludo.NewCard("games room"),
		"THIS",
		"charlie",
	)
	q.SetAnswer(cludo.WhereAnswer)
	game.DoTurn(q)

	fmt.Println(game)
}
