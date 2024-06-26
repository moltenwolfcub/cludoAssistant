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

	q.Cards[0].SetFound(nil, true)
	q.Cards[1].SetFound(nil, true)
	q.Cards[2].SetFound(nil, true)

	if !q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Couldn't find solution when there was one present.")
	}
}

func TestHasKnownSolutionWithoutSolution(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Cards[0].SetFound(nil, true)
	q.Cards[1].SetFound(nil, true)

	if q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Found a solution when there were multiple options.")
	}
}

func TestHasKnownSolutionWithoutOptions(t *testing.T) {
	q := GenSampleQuestionCategory()

	q.Cards[0].SetFound(nil, true)
	q.Cards[1].SetFound(nil, true)
	q.Cards[2].SetFound(nil, true)
	q.Cards[3].SetFound(nil, true)

	if q.HasKnownSolution() {
		t.Error("Question.HasKnownSolution() Found a solution when there were no options left.")
	}
}

func GenSampleGame() (game cluedo.Game, a, b, c *cluedo.Player) {
	a = cluedo.NewPlayer("alice", 0)
	b = cluedo.NewPlayer("bob", 0)
	c = cluedo.NewPlayer("charlie", 0)

	game = cluedo.NewDefaultGame(a, b, c)
	return
}

func TestEnsureValidQuestionWithValid(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("plum"),
		cluedo.NewCard("dagger"),
		cluedo.NewCard("study"),
		game.Me,
		alice,
	)

	if !game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question wasn't valid when it should've been")
	}
}

func TestEnsureValidQuestionWithInvalidQuestionComponents(t *testing.T) {
	game, alice, _, _ := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("eva smith"),
		cluedo.NewCard("bleach"),
		cluedo.NewCard("infermary"),
		game.Me,
		alice,
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when it wasn't made of elements from the game")
	}
}

func TestEnsureValidQuestionWithInvalidPlayers(t *testing.T) {
	game, _, _, _ := GenSampleGame()

	question := cluedo.NewQuestion(
		cluedo.NewCard("mustard"),
		cluedo.NewCard("rope"),
		cluedo.NewCard("kitchen"),
		cluedo.NewPlayer("eric", 0),
		cluedo.NewPlayer("inspector goole", 0),
	)

	if game.EnsureValidQuestion(question) {
		t.Error("Game.EnsureValidQuestion() Question was deemed valid when the players weren't ones in the game")
	}
}
