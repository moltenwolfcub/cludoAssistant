package cludo_test

import (
	"testing"

	"github.com/moltenwolfcub/cludoAssistant/cludo"
)

func GenSampleQuestionCategory() cludo.QuestionCategory {
	return cludo.NewQuestion(
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
