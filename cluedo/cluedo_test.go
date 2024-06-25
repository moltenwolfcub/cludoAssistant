package cluedo_test

import (
	"testing"

	"github.com/moltenwolfcub/cluedoAssistant/cluedo"
)

func GenSampleQuestionCategory() cluedo.QuestionCategory {
	return cluedo.NewQuestionCategory(
		cluedo.NewCard("Zero"),
		cluedo.NewCard("One"),
		cluedo.NewCard("two"),
		cluedo.NewCard("three"),
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

func GenSampleGame() cluedo.Game {
	players := []string{
		"alice",
		"bob",
		"charlie",
	}
	return cluedo.NewDefaultGame(players)
}

func TestEnsureValidQuestionWithValid(t *testing.T) {
	game := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("plum"),
		cluedo.NewCard("dagger"),
		cluedo.NewCard("study"),
		cluedo.Player("THIS"),
		cluedo.Player("alice"),
	)

	if !game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question wasn't valid when it should've been")
	}
}

func TestEnsureValidQuestionWithInvalidQuestionComponents(t *testing.T) {
	game := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("eva smith"),
		cluedo.NewCard("bleach"),
		cluedo.NewCard("the factory"),
		cluedo.Player("THIS"),
		cluedo.Player("alice"),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when it wasn't made of elements from the game")
	}
}

func TestEnsureValidQuestionWithInvalidPlayers(t *testing.T) {
	game := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("mustard"),
		cluedo.NewCard("rope"),
		cluedo.NewCard("kitchen"),
		cluedo.Player("eric"),
		cluedo.Player("inspector goole"),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when the players weren't ones in the game")
	}
}
