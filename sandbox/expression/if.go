package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/code_block"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type If struct {
	condition concept.Index
	judgment  *code_block.CodeBlock
	primary   *code_block.CodeBlock
	secondary *code_block.CodeBlock
}

func (f *If) ToString(prefix string) string {
	judgmentToString := ""

	if f.judgment.Size() != 0 {
		judgmentToString = fmt.Sprintf(" %v", f.judgment.ToString(prefix))
	}

	primaryToString := fmt.Sprintf("%vif (%v%v) %v", prefix, f.condition.ToString(prefix), judgmentToString, f.primary.ToString(prefix))
	if f.secondary.Size() == 0 {
		return primaryToString
	}
	return fmt.Sprintf("%v else %v", primaryToString, f.secondary.ToString(prefix))
}

func (f *If) Exec(parent concept.Closure) concept.Interrupt {

	if f.condition == nil {
		return interrupt.NewException("system error", "No condition for judgment.")
	}

	judgmentSpace, suspend := f.judgment.Exec(parent, false, nil)
	defer judgmentSpace.Clear()
	if !nl_interface.IsNil(suspend) {
		return suspend
	}

	preCondition, suspend := f.condition.Get(judgmentSpace)
	if !nl_interface.IsNil(suspend) {
		return suspend
	}

	condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return interrupt.NewException("type error", "Only bool can be judged.")
	}

	var active *code_block.CodeBlock
	if condition.Value() {
		active = f.primary
	} else {
		active = f.secondary
	}

	space, suspend := active.Exec(judgmentSpace, true, nil)
	defer space.Clear()
	parent.MergeReturn(judgmentSpace)
	return suspend
}

func (f *If) SetCondition(condition concept.Index) {
	f.condition = condition
}

func (f *If) AddJudgmentStep(step concept.Expression) {
	f.judgment.AddStep(step)
}

func (f *If) AddPrimaryStep(step concept.Expression) {
	f.primary.AddStep(step)
}

func (f *If) AddSecondaryStep(step concept.Expression) {
	f.secondary.AddStep(step)
}

func NewIf() *If {
	return &If{
		judgment:  code_block.NewCodeBlock(),
		primary:   code_block.NewCodeBlock(),
		secondary: code_block.NewCodeBlock(),
	}
}
