package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
)

type FunctionEnd struct {
	*adaptor.ExpressionIndex
}

var (
	FunctionEndLanguageSeeds = map[string]func(string, *FunctionEnd) string{}
)

func (f *FunctionEnd) ToLanguage(language string) string {
	seed := FunctionEndLanguageSeeds[language]
	if seed == nil {
		return f.ToString("")
	}
	return seed(language, f)
}

func (a *FunctionEnd) ToString(prefix string) string {
	return fmt.Sprintf("end")
}

func (a *FunctionEnd) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return nil, interrupt.NewEnd()
}

func NewFunctionEnd() *FunctionEnd {
	back := &FunctionEnd{}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
