package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
)

type Return struct {
	*adaptor.ExpressionIndex
	key    concept.String
	result concept.Index
}

func (a *Return) Key() concept.String {
	return a.key
}

func (a *Return) ToString(prefix string) string {
	return fmt.Sprintf("return[%v] %v", a.key.ToString(prefix), a.result.ToString(prefix))
}

func (a *Return) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	result, suspend := a.result.Get(space)

	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}
	space.SetReturn(a.key, result)
	return result, nil
}

func NewReturn(key concept.String, result concept.Index) *Return {
	back := &Return{
		key:    key,
		result: result,
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
