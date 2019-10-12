package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/closure"
	"github.com/TingerSure/natural_language/sandbox/code_block"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type If struct {
	condition concept.Index
	primary   *code_block.CodeBlock
	secondary *code_block.CodeBlock
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
		return nil, interrupt.NewException("system error", "No condition for judgment.")
	}
	initSpace := closure.NewClosure(parent)
	defer initSpace.Clear()
	defer parent.MergeReturn(initSpace)

	preCondition, suspend := f.condition.Get(initSpace)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return nil, interrupt.NewException("type error", "Only bool can be judged.")
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
	return &If{
		primary:   code_block.NewCodeBlock(),
		secondary: code_block.NewCodeBlock(),
	}
}
