package sandbox

type If struct {
	condition Index
	judgment  *CodeBlock
	primary   *CodeBlock
	secondary *CodeBlock
}

func (f *If) Exec(parent *Closure) Interrupt {

	if f.condition == nil {
		return NewException("system error", "No condition for judgment.")
	}

	judgmentSpace, suspend := f.judgment.Exec(parent, false, nil)
	defer judgmentSpace.Clear()
	if suspend != nil {
		return suspend
	}

	preCondition, suspend := f.condition.Get(judgmentSpace)
	if suspend != nil {
		return suspend
	}

	condition, yes := VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return NewException("type error", "Only bool can be judged.")
	}

	var active *CodeBlock
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
