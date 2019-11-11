package tree

type PriorityRule struct {
	match   func(Phrase, Phrase) bool
	chooser func(Phrase, Phrase) int
}

func (p *PriorityRule) Match(left, right Phrase) bool {
	return p.match(left, right)
}

func (p *PriorityRule) Choose(left, right Phrase) int {
	return p.chooser(left, right)
}

func NewPriorityRule(
	match func(Phrase, Phrase) bool,
	chooser func(Phrase, Phrase) int,
) *PriorityRule {
	return &PriorityRule{
		match:   match,
		chooser: chooser,
	}
}
