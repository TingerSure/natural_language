package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type Return struct {
	key    string
	result concept.Index
}

func (a *Return) ToString(prefix string) string {
	return fmt.Sprintf("return[%v] %v", a.key, a.result.ToString(prefix))
}

func (a *Return) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	result, suspend := a.result.Get(space)

	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	space.SetReturn(a.key, result)
	return result, nil
}

func NewReturn(key string, result concept.Index) *Return {
	return &Return{
		key:    key,
		result: result,
	}
}
