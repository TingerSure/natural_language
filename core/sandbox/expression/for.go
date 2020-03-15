package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	expressionForDefaultCondition = index.NewConstIndex(variable.NewBool(true))
)

type For struct {
	*adaptor.ExpressionIndex
	tag       string
	condition concept.Index
	init      *code_block.CodeBlock
	end       *code_block.CodeBlock
	body      *code_block.CodeBlock
}

func (f *For) SubCodeBlockIterate(onIndex func(concept.Index) bool) bool {
	return f.init.Iterate(onIndex) || f.end.Iterate(onIndex) || f.body.Iterate(onIndex)
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
func (f *For) Body() *code_block.CodeBlock {
	return f.body
}
func (f *For) Init() *code_block.CodeBlock {
	return f.init
}
func (f *For) End() *code_block.CodeBlock {
	return f.end
}

func NewFor() *For {
	param := &code_block.CodeBlockParam{
		StringCreator: func(value string) concept.String {
			return variable.NewString(value)
		},
	}
	back := &For{
		tag:  "",
		init: code_block.NewCodeBlock(param),
		end:  code_block.NewCodeBlock(param),
		body: code_block.NewCodeBlock(param),
	}
	back.ExpressionIndex = adaptor.NewExpressionIndex(back.Exec)
	return back
}
