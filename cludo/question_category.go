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

func NewQuestion(options ...Option) QuestionCategory {
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
