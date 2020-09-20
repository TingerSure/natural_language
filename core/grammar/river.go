package grammar

import (
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type River struct {
	lake *Lake
	flow *lexer.Flow
}

func (r *River) ToString() string {
	space := ""
	for _, phrase := range r.lake.PeekAll() {
		space += phrase.ToString()
	}
	return space
}

func (r *River) IsActive() bool {
	return r.lake.IsSingle() && r.flow.IsEnd()
}

func (r *River) GetLake() *Lake {
	return r.lake
}

func (r *River) GetFlow() *lexer.Flow {
	return r.flow
}

func (r *River) Step(section *Section, reach *Reach, dam *Dam) ([]*River, error) {
	var err error
	tributaries := []*River{}
	subStructs := []*River{}
	subVocabularies := []*River{}
	reach.Check(r.lake, func(twig *tree.StructRule) {
		tributary := r.Copy()
		phrase := twig.Create(tributary.lake.PeekAll())
		tributary.lake.PopMultiple(twig.Size())
		tributary.lake.Push(phrase)
		subStructs = append(subStructs, tributary)
	})
	err = section.Check(r.flow, func(rule *tree.VocabularyRule) {
		tributary := r.Copy()
		tributary.lake.Push(rule.Create(tributary.flow.Next()))
		subVocabularies = append(subVocabularies, tributary)
	})
	if err != nil {
		return nil, err
	}
	if len(subStructs) == 0 && len(subVocabularies) == 0 {
		tributaries = append(tributaries, r)
		return tributaries, nil
	}
	for _, subStruct := range subStructs {
		subs, err := subStruct.Step(section, reach, dam)
		if err != nil {
			return nil, err
		}
		subs = dam.Filter(subs)
		tributaries = append(tributaries, subs...)
	}

	for _, subVocabulary := range subVocabularies {
		subs, err := subVocabulary.Step(section, reach, dam)
		if err != nil {
			return nil, err
		}
		tributaries = append(tributaries, subs...)
	}
	return tributaries, nil
}

func (r *River) Copy() *River {
	return NewRiver(r.lake.Copy(), r.flow.Copy())
}

func NewRiver(lake *Lake, flow *lexer.Flow) *River {
	return &River{
		lake: lake,
		flow: flow,
	}
}
