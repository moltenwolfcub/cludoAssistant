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

type Question struct {
	Options []Option
}

func NewQuestion(options ...Option) Question {
	return Question{
		Options: options,
	}
}

func (q Question) HasKnownSolution() bool {
	return false
}
