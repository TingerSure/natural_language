package grammar

import (
	"github.com/TingerSure/natural_language/core/lexer"
)

type Grammar struct {
	reach   *Reach
	section *Section
	dam     *Dam
}

func (g *Grammar) GetDam() *Dam {
	return g.dam
}

func (g *Grammar) GetSection() *Section {
	return g.section
}

func (g *Grammar) GetReach() *Reach {
	return g.reach
}

func (l *Grammar) Instances(flow *lexer.Flow) (*Valley, error) {
	flow.Reset()
	valley := NewValley()
	err := valley.Step(flow, l.section, l.reach, l.dam)
	if err != nil {
		return nil, err
	}

	return valley, nil
}

func (l *Grammar) init() *Grammar {
	return l
}

func NewGrammar() *Grammar {
	return (&Grammar{
		dam:     NewDam(),
		section: NewSection(),
		reach:   NewReach(),
	}).init()
}
