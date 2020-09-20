package grammar

import (
	"errors"
	"fmt"
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

func (r *River) structCheck(twigs []*tree.StructRule, onStruct func(*tree.StructRule)) {
	if r.lake.IsEmpty() {
		return
	}
	for _, twig := range twigs {
		if r.lake.Len() < twig.Size() {
			continue
		}
		if twig.Match(r.lake.PeekAll()) {
			onStruct(twig)
		}
	}
}

func (r *River) vocabularyCheck(leaves []*tree.VocabularyRule, onVocabulary func(*tree.VocabularyRule)) error {
	if r.flow.IsEnd() {
		return nil
	}
	word := r.flow.Peek()
	count := 0
	for _, leaf := range leaves {
		if leaf.Match(word) {
			onVocabulary(leaf)
			count++
		}
	}
	if count == 0 {
		return errors.New(fmt.Sprintf("This vocabulary has no rules to parse! ( %v )", word.ToString()))
	}
	return nil
}

func (r *River) Step(leaves []*tree.VocabularyRule, twigs []*tree.StructRule, dam *Dam) ([]*River, error) {
	var err error
	tributaries := []*River{}
	subStructs := []*River{}
	subVocabularies := []*River{}
	r.structCheck(twigs, func(twig *tree.StructRule) {
		tributary := r.Copy()
		phrase := twig.Create(tributary.lake.PeekAll())
		tributary.lake.PopMultiple(twig.Size())
		tributary.lake.Push(phrase)
		subStructs = append(subStructs, tributary)
	})
	err = r.vocabularyCheck(leaves, func(leaf *tree.VocabularyRule) {
		tributary := r.Copy()
		tributary.lake.Push(leaf.Create(tributary.flow.Next()))
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
		subs, err := subStruct.Step(leaves, twigs, dam)
		if err != nil {
			return nil, err
		}
		subs = dam.Filter(subs)
		tributaries = append(tributaries, subs...)
	}

	for _, subVocabulary := range subVocabularies {
		subs, err := subVocabulary.Step(leaves, twigs, dam)
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
