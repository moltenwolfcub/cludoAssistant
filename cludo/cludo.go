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

type Question struct {
	options []Option
}

func NewQuestion(options ...Option) Question {
	return Question{
		options: options,
	}
}
