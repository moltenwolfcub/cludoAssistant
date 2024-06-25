package cludo_test

import (
	"testing"

	"github.com/moltenwolfcub/cludoAssistant/cludo"
)

func GenSampleQuestionCategory() cludo.QuestionCategory {
	return cludo.NewQuestionCategory(
		cludo.NewCard("Zero"),
		cludo.NewCard("One"),
		cludo.NewCard("two"),
		cludo.NewCard("three"),
	)
}

func TestHasKnownSolutionWithSolution(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Cards[0].SetFound("", true)
	q.Cards[1].SetFound("", true)
	q.Cards[2].SetFound("", true)

	if !q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Couldn't find solution when there was one present.")
	}
}

func TestHasKnownSolutionWithoutSolution(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Cards[0].SetFound("", true)
	q.Cards[1].SetFound("", true)

	if q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Found a solution when there were multiple options.")
	}
}

func TestHasKnownSolutionWithoutOptions(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Cards[0].SetFound("", true)
	q.Cards[1].SetFound("", true)
	q.Cards[2].SetFound("", true)
	q.Cards[3].SetFound("", true)

	if q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Found a solution when there were no options left.")
	}
}

func GenSampleGame() cludo.Game {
	players := []string{
		"alice",
		"bob",
		"charlie",
	}
	return cludo.NewDefaultGame(players)
}

func TestEnsureValidQuestionWithValid(t *testing.T) {
	game := GenSampleGame()

	question := cludo.NewQuestion(
		cludo.NewCard("plum"),
		cludo.NewCard("dagger"),
		cludo.NewCard("study"),
		cludo.Player("THIS"),
		cludo.Player("alice"),
	)

	if !game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question wasn't valid when it should've been")
	}
}

func TestEnsureValidQuestionWithInvalidQuestionComponents(t *testing.T) {
	game := GenSampleGame()

	question := cludo.NewQuestion(
		cludo.NewCard("eva smith"),
		cludo.NewCard("bleach"),
		cludo.NewCard("the factory"),
		cludo.Player("THIS"),
		cludo.Player("alice"),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when it wasn't made of elements from the game")
	}
}

func TestEnsureValidQuestionWithInvalidPlayers(t *testing.T) {
	game := GenSampleGame()

	question := cludo.NewQuestion(
		cludo.NewCard("mustard"),
		cludo.NewCard("rope"),
		cludo.NewCard("kitchen"),
		cludo.Player("eric"),
		cludo.Player("inspector goole"),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when the players weren't ones in the game")
	}
}
