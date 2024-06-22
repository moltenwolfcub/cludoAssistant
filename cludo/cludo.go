package cludo

type Option struct {
	name  string
	found bool
}

func NewOption(name string) Option {
	return Option{
		name:  name,
		found: false,
	}
}

func (o *Option) SetFound() {
	o.found = true
}

type QuestionCategory struct {
	Options []Option
}

func NewQuestionCategory(options ...Option) QuestionCategory {
	return QuestionCategory{
		Options: options,
	}
}

func (q QuestionCategory) HasKnownSolution() bool {
	available := []Option{}
	for _, o := range q.Options {
		if !o.found {
			available = append(available, o)
		}
	}
	return len(available) == 1
}

type Player string

type Game struct {
	who   QuestionCategory
	what  QuestionCategory
	where QuestionCategory

	otherPlayers []Player
}

func NewDefaultGame(otherPlayers []string) Game {
	g := Game{
		who: NewQuestionCategory(
			NewOption("green"),
			NewOption("mustard"),
			NewOption("peacock"),
			NewOption("plum"),
			NewOption("scarlet"),
			NewOption("white"),
		),
		what: NewQuestionCategory(
			NewOption("wrench"),
			NewOption("candlestick"),
			NewOption("dagger"),
			NewOption("pistol"),
			NewOption("lead pipe"),
			NewOption("rope"),
		),
		where: NewQuestionCategory(
			NewOption("bathroom"),
			NewOption("study"),
			NewOption("dining room"),
			NewOption("games room"),
			NewOption("garage"),
			NewOption("bedroom"),
			NewOption("living room"),
			NewOption("kitchen"),
			NewOption("courtyard"),
		),
	}

	for _, p := range otherPlayers {
		g.otherPlayers = append(g.otherPlayers, Player(p))
	}

	return g
}
