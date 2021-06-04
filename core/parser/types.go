package parser

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/tree"
)

type Types struct {
	values map[string]*tree.PhraseType
}

func NewTypes() *Types {
	return &Types{
		values: map[string]*tree.PhraseType{},
	}
}

func (t *Types) Package(wanted string, given string, from tree.Phrase) tree.Phrase {
	isMatched, rules := t.match(wanted, given)
	if !isMatched {
		panic(fmt.Sprintf("PhraseType Error, wanted (%v) and given (%v) are unmatch.", wanted, given))
	}

	for index := len(rules) - 1; index >= 0; index-- {
		if rules[index] != nil {
			from = rules[index].Pack(from)
		}
	}
	return from
}

func (t *Types) Match(wanted string, given string) bool {
	isMatch, _ := t.match(wanted, given)
	return isMatch
}

func (t *Types) match(wanted string, given string) (bool, []*tree.PackageRule) {
	if wanted == given {
		return true, nil
	}

	for _, givenParent := range t.values[given].Parents() {
		isMatch, rules := t.match(wanted, givenParent.Types)
		if isMatch {
			return true, append(rules, givenParent.Rule)
		}
	}

	return false, nil
}

func (t *Types) IterateParentsDistinct(wanted string, onParent func(string) bool) bool {
	record := map[string]bool{}
	for _, parent := range t.values[wanted].Parents() {
		if !record[parent.Types] {
			if t.IterateParentsDistinct(parent.Types, func(grandParent string) bool {

				if !record[grandParent] {
					record[grandParent] = true
					return onParent(grandParent)
				}
				return false
			}) {
				return true
			}
			record[parent.Types] = true
			if onParent(parent.Types) {
				return true
			}
		}
	}
	return false
}

func (t *Types) GetTypes(name string) *tree.PhraseType {
	return t.values[name]
}

func (t *Types) AddTypes(value *tree.PhraseType) {
	t.values[value.Name()] = value
}

func (t *Types) RemoveTypes(need func(types *tree.PhraseType) bool) {
	for key, value := range t.values {
		if need(value) {
			delete(t.values, key)
		}
	}
}
