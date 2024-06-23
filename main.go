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
		[]*cludo.Option{
			cludo.NewOption("peacock"),
			cludo.NewOption("white"),
			cludo.NewOption("rope"),
			cludo.NewOption("bathroom"),
		},
	)

	q := cludo.NewQuestion(
		cludo.NewOption("scarlet"),
		cludo.NewOption("dagger"),
		cludo.NewOption("study"),
		"THIS",
		"alice",
	)
	q.SetAnswer(cludo.WhoAnswer)
	game.DoTurn(q)

	q = cludo.NewQuestion(
		cludo.NewOption("green"),
		cludo.NewOption("lead pipe"),
		cludo.NewOption("games room"),
		"THIS",
		"bob",
	)
	q.SetAnswer(cludo.WhoAnswer)
	game.DoTurn(q)

	fmt.Println(game)
}
