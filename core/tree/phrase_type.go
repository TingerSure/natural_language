package tree

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
)

type PhraseType struct {
	parents []*PhraseType
	name    string
}

func (wanted *PhraseType) Match(given *PhraseType) bool {
	if nl_interface.IsNil(given) {
		return false
	}

	if wanted.Equal(given) {
		return true
	}

	for _, givenParent := range given.parents {
		if wanted.Match(givenParent) {
			return true
		}
	}

	return false
}

func (wanted *PhraseType) Equal(given *PhraseType) bool {
	return wanted == given
}

func (p *PhraseType) Name() string {
	return p.name
}

func (p *PhraseType) Parents() []*PhraseType {
	return p.parents
}

func NewPhraseType(name string, parents []*PhraseType) *PhraseType {
	if parents == nil {
		parents = []*PhraseType{}
	}
	return &PhraseType{
		parents: parents,
		name:    name,
	}
}
