package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type If struct {
	*adaptor.ExpressionIndex
	condition concept.Index
	primary   *code_block.CodeBlock
	secondary *code_block.CodeBlock
}

func (f *If) SubCodeBlockIterate(onIndex func(concept.Index) bool) bool {
	return f.primary.Iterate(onIndex) || f.secondary.Iterate(onIndex)
}

func (f *If) ToString(prefix string) string {
	primaryToString := fmt.Sprintf("if (%v) %v", f.condition.ToString(prefix), f.primary.ToString(prefix))
	if f.secondary.Size() == 0 {
		return primaryToString
	}
	return fmt.Sprintf("%v else %v", primaryToString, f.secondary.ToString(prefix))
}

func (f *If) Exec(parent concept.Closure) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		return nil, interrupt.NewException(variable.NewString("system error"), variable.NewString("No condition for judgment."))
	}
	initSpace := closure.NewClosure(parent, &closure.ClosureParam{
		StringCreator: func(value string) concept.String {
			return variable.NewString(value)
		},
	})
	defer initSpace.Clear()
	defer parent.MergeReturn(initSpace)

	preCondition, suspend := f.condition.Get(initSpace)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return nil, interrupt.NewException(variable.NewString("type error"), variable.NewString("Only bool can be judged."))
	}

	var active *code_block.CodeBlock
	if condition.Value() {
		active = f.primary
	} else {
		active = f.secondary
	}

	space, suspend := active.Exec(initSpace, true, nil)
	defer space.Clear()
	return nil, suspend
}

func (f *If) SetCondition(condition concept.Index) {
	f.condition = condition
}

func (f *If) Primary() *code_block.CodeBlock {
	return f.primary
}

func (f *If) Secondary() *code_block.CodeBlock {
	return f.secondary
}

func NewIf() *If {
	param := &code_block.CodeBlockParam{
		StringCreator: func(value string) concept.String {
			return variable.NewString(value)
		},
	}
	back := &If{
		primary:   code_block.NewCodeBlock(param),
		secondary: code_block.NewCodeBlock(param),
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
