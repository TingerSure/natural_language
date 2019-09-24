package sandbox

import (
	"errors"
)

type If struct {
	condition Index
	judgment  *CodeBlock
	primary   *CodeBlock
	secondary *CodeBlock
}

func (f *If) Exec(parent *Closure) (bool, error) {

	if f.condition == nil {
		return false, errors.New("No condition for judgment.")
	}

	judgmentSpace, keep, err := f.judgment.Exec(parent, false, nil)
	if err != nil {
		return false, err
	}
	defer judgmentSpace.Clear()
	if !keep {
		return false, nil
	}

	preCondition, err := f.condition.Get(judgmentSpace)
	if err != nil {
		return false, err
	}

	condition, yes := VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return false, errors.New("Only bool can be judged.")
	}

	var active *CodeBlock
	if condition.Value() {
		active = f.primary
	} else {
		active = f.secondary
	}

	space, keep, err := active.Exec(judgmentSpace, true, nil)
	if err != nil {
		return false, err
	}
	defer space.Clear()
	parent.MergeReturn(judgmentSpace)
	return keep, nil
}

func (f *If) SetCondition(condition Index) {
	f.condition = condition
}

func (f *If) AddJudgmentStep(step Expression) {
	f.judgment.AddStep(step)
}

func (f *If) AddPrimaryStep(step Expression) {
	f.primary.AddStep(step)
}

func (f *If) AddSecondaryStep(step Expression) {
	f.secondary.AddStep(step)
}

func NewIf() *If {
	return &If{
		judgment:  NewCodeBlock(),
		primary:   NewCodeBlock(),
		secondary: NewCodeBlock(),
	}
}
