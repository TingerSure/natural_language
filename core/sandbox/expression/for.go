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

type ForSeed interface {
	ToLanguage(string, *For) string
	GetDefaultCondition() concept.Index
	GetDefaultTag() concept.String
	NewException(string, string) concept.Exception
	NewNull() concept.Null
}

type For struct {
	*adaptor.ExpressionIndex
	tag       concept.String
	condition concept.Index
	init      *code_block.CodeBlock
	end       *code_block.CodeBlock
	body      *code_block.CodeBlock
	seed      ForSeed
}

func (f *For) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (f *For) SubCodeBlockIterate(onIndex func(concept.Index) bool) bool {
	return f.init.Iterate(onIndex) || f.end.Iterate(onIndex) || f.body.Iterate(onIndex)
}

func (f *For) ToString(prefix string) string {
	return fmt.Sprintf("for (%v; %v; %v) %v", f.init.ToStringSimplify(prefix), f.condition.ToString(prefix), f.end.ToStringSimplify(prefix), f.body.ToString(prefix))
}

func (e *For) Anticipate(space concept.Closure) concept.Variable {
	return e.seed.NewNull()
}

func (f *For) Exec(parent concept.Closure) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		f.condition = f.seed.GetDefaultCondition()
	}

	initSpace, suspend := f.init.Exec(parent, false, nil)
	defer initSpace.Clear()
	defer parent.MergeReturn(initSpace)
	if !nl_interface.IsNil(suspend) {
		return f.seed.NewNull(), suspend
	}

body:
	for {
		preCondition, suspend := f.condition.Get(initSpace)
		if !nl_interface.IsNil(suspend) {
			return f.seed.NewNull(), suspend
		}

		condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
		if !yes {
			return f.seed.NewNull(), f.seed.NewException("type error", "Only bool can be judged.")
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
					return f.seed.NewNull(), f.seed.NewException("system panic", fmt.Sprintf("BreakInterruptType does not mean a Break anymore.\n%+v", suspend))
				}
				if !f.IsMyTag(breaks.Tag()) {
					return f.seed.NewNull(), suspend
				}
				break body
			case interrupt.ContinueInterruptType:
				continues, yes := interrupt.InterruptFamilyInstance.IsContinue(suspend)
				if !yes {
					return f.seed.NewNull(), f.seed.NewException("system panic", fmt.Sprintf("ContinueInterruptType does not mean a Continue anymore.\n%+v", suspend))
				}
				if !f.IsMyTag(continues.Tag()) {
					return f.seed.NewNull(), suspend
				}
			default:
				return f.seed.NewNull(), suspend
			}
		}
		endSpace, suspend := f.end.Exec(initSpace, true, nil)
		defer endSpace.Clear()
		if !nl_interface.IsNil(suspend) {
			return f.seed.NewNull(), suspend
		}
	}

	return f.seed.NewNull(), nil
}

func (f *For) SetTag(tag concept.String) {
	f.tag = tag
}
func (f *For) Tag() concept.String {
	return f.tag
}
func (f *For) IsMyTag(tag concept.String) bool {
	if tag.Equal(f.seed.GetDefaultTag()) ||
		tag.Equal(f.tag) ||
		tag.EqualLanguage(f.tag) {
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

type ForCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	StringCreator          func(string) concept.String
	BoolCreator            func(bool) *variable.Bool
	CodeBlockCreator       func() *code_block.CodeBlock
	ConstIndexCreator      func(concept.Variable) *index.ConstIndex
	ExpressionIndexCreator func(func(concept.Closure) (concept.Variable, concept.Interrupt)) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type ForCreator struct {
	Seeds            map[string]func(string, *For) string
	param            *ForCreatorParam
	defaultCondition concept.Index
	defaultTag       concept.String
}

func (s *ForCreator) New() *For {
	back := &For{
		tag:  s.defaultTag,
		init: s.param.CodeBlockCreator(),
		end:  s.param.CodeBlockCreator(),
		body: s.param.CodeBlockCreator(),
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back.Exec)
	return back
}

func (s *ForCreator) ToLanguage(language string, instance *For) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *ForCreator) GetDefaultCondition() concept.Index {
	return s.defaultCondition
}

func (s *ForCreator) GetDefaultTag() concept.String {
	return s.defaultTag
}

func (s *ForCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *ForCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewForCreator(param *ForCreatorParam) *ForCreator {
	return &ForCreator{
		Seeds:            map[string]func(string, *For) string{},
		param:            param,
		defaultCondition: param.ConstIndexCreator(param.BoolCreator(true)),
		defaultTag:       param.StringCreator(""),
	}
}
