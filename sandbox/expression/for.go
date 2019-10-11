package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/library/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/code_block"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/sandbox/interrupt"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	expressionForDefaultCondition = index.NewConstIndex(variable.NewBool(true))
)

type For struct {
	tag       string
	condition concept.Index
	init      *code_block.CodeBlock
	end       *code_block.CodeBlock
	body      *code_block.CodeBlock
}

func (f *For) ToString(prefix string) string {
	return fmt.Sprintf("for (%v; %v; %v) %v", f.init.ToStringSimplify(prefix), f.condition.ToString(prefix), f.end.ToStringSimplify(prefix), f.body.ToString(prefix))
}

func (f *For) Exec(parent concept.Closure) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		f.condition = expressionForDefaultCondition
	}

	initSpace, suspend := f.init.Exec(parent, false, nil)
	defer initSpace.Clear()
	defer parent.MergeReturn(initSpace)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

body:
	for {
		preCondition, suspend := f.condition.Get(initSpace)
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}

		condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
		if !yes {
			return nil, interrupt.NewException("type error", "Only bool can be judged.")
		}

		if !condition.Value() {
			break body
		}

		space, suspend := f.body.Exec(initSpace, true, nil)
		defer space.Clear()
		if !nl_interface.IsNil(suspend) {
			switch suspend.InterruptType() {
			case interrupt.BreakInterruptType:
				breaks, yes := interrupt.InterruptFamilyInstance.IsBreak(suspend)
				if !yes {
					return nil, interrupt.NewException("system panic", fmt.Sprintf("BreakInterruptType does not mean a Break anymore.\n%+v", suspend))
				}
				if !f.IsMyTag(breaks.Tag()) {
					return nil, suspend
				}
				break body
			case interrupt.ContinueInterruptType:
				continues, yes := interrupt.InterruptFamilyInstance.IsContinue(suspend)
				if !yes {
					return nil, interrupt.NewException("system panic", fmt.Sprintf("ContinueInterruptType does not mean a Continue anymore.\n%+v", suspend))
				}
				if !f.IsMyTag(continues.Tag()) {
					return nil, suspend
				}
			default:
				return nil, suspend
			}
		}
		endSpace, suspend := f.end.Exec(initSpace, true, nil)
		defer endSpace.Clear()
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
		}
	}

	return nil, nil
}

func (f *For) SetTag(tag string) {
	f.tag = tag
}
func (f *For) Tag() string {
	return f.tag
}
func (f *For) IsMyTag(tag string) bool {
	if tag == "" || tag == f.tag {
		return true
	}
	return false
}

func (f *For) SetCondition(condition concept.Index) {
	f.condition = condition
}

func (f *For) AddBodyStep(step concept.Expression) {
	f.body.AddStep(step)
}

func (f *For) AddInitStep(step concept.Expression) {
	f.init.AddStep(step)
}

func (f *For) AddEndStep(step concept.Expression) {
	f.end.AddStep(step)
}

func NewFor() *For {
	return &For{
		tag:  "",
		init: code_block.NewCodeBlock(),
		end:  code_block.NewCodeBlock(),
		body: code_block.NewCodeBlock(),
	}
}
