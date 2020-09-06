package grammar

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/tree"
)

type Valley struct {
	values []*River
}

func (v *Valley) Size() int {
	return len(v.values)
}

func (v *Valley) Step(flow *lexer.Flow, leaves []*tree.VocabularyRule, twigs []*tree.StructRule) error {
	wait := NewCollector()
	river := NewRiver(wait, flow)
	bases, err := river.Step(leaves, twigs)
	if err != nil {
		return err
	}
	for _, base := range bases {
		v.AddRiver(base)
	}
	return nil
}

func (v *Valley) AddRiver(value *River) {
	v.values = append(v.values, value)
}

func (v *Valley) Iterate(onRiver func(river *River) bool) bool {
	for _, value := range v.values {
		if onRiver(value) {
			return true
		}
	}
	return false
}

func (v *Valley) Filter() (*Valley, error) {
	if v.Size() == 0 {
		return nil, errors.New("This is an empty sentence!")
	}
	actives := NewValley()
	var min *River = nil

	v.Iterate(func(input *River) bool {
		if input.IsActive() {
			actives.AddRiver(input)
			return false
		}
		if min == nil {
			min = input
		}
		if input.GetWait().Len() < min.GetWait().Len() {
			min = input
		}
		return false
	})
	if actives.Size() == 0 {
		return nil, errors.New(fmt.Sprintf("There is an unknown rule in this sentence!\n%v", min.ToString()))
	}
	return actives, nil
}

func NewValley() *Valley {
	return &Valley{}
}
