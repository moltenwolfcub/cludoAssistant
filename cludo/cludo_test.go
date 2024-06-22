package cludo_test

import (
	"testing"

	"github.com/moltenwolfcub/cludoAssistant/cludo"
)

func GenSampleQuestionCategory() cludo.QuestionCategory {
	return cludo.NewQuestionCategory(
		cludo.NewOption("Zero"),
		cludo.NewOption("One"),
		cludo.NewOption("two"),
		cludo.NewOption("three"),
	)
}

func TestHasKnownSolutionWithSolution(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Options[0].SetFound()
	q.Options[1].SetFound()
	q.Options[2].SetFound()

	if !q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Couldn't find solution when there was one present.")
	}
}

func TestHasKnownSolutionWithoutSolution(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Options[0].SetFound()
	q.Options[1].SetFound()

	if q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Found a solution when there were multiple options.")
	}
}

func TestHasKnownSolutionWithoutOptions(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Options[0].SetFound()
	q.Options[1].SetFound()
	q.Options[2].SetFound()
	q.Options[3].SetFound()

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
		cludo.NewOption("plum"),
		cludo.NewOption("dagger"),
		cludo.NewOption("study"),
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
		cludo.NewOption("eva smith"),
		cludo.NewOption("bleach"),
		cludo.NewOption("the factory"),
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
		cludo.NewOption("mustard"),
		cludo.NewOption("rope"),
		cludo.NewOption("kitchen"),
		cludo.Player("eric"),
		cludo.Player("inspector goole"),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when the players weren't ones in the game")
	}
}
