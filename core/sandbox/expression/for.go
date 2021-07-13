package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/interrupt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type ForSeed interface {
	ToLanguage(string, concept.Pool, *For) (string, concept.Exception)
	GetDefaultCondition() concept.Pipe
	GetDefaultTag() concept.String
	NewException(string, string) concept.Exception
	NewNull() concept.Null
}

type For struct {
	*adaptor.ExpressionIndex
	tag       concept.String
	condition concept.Pipe
	init      concept.CodeBlock
	end       concept.CodeBlock
	body      concept.CodeBlock
	line      concept.Line
	seed      ForSeed
}

func (f *For) SetLine(line concept.Line) {
	f.line = line
}

func (f *For) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (f *For) ToString(prefix string) string {
	back := ""
	if f.condition != f.seed.GetDefaultCondition() {
		back = f.condition.ToString(prefix)
	}
	if f.init.Size() != 0 || f.end.Size() != 0 {
		back = fmt.Sprintf("%v; %v; %v", f.init.ToStringSimplify(prefix), back, f.end.ToStringSimplify(prefix))
	}
	back = fmt.Sprintf("for (%v)", back)
	if f.tag != f.seed.GetDefaultTag() {
		back = fmt.Sprintf("%v : %v", f.tag.Value(), back)
	}
	return fmt.Sprintf("%v %v", back, f.body.ToString(prefix))
}

func (f *For) Exec(parent concept.Pool) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		f.condition = f.seed.GetDefaultCondition()
	}

	initSpace, suspend := f.init.ExecWithInit(parent, nil)
	defer initSpace.Clear()
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
			return nil, f.seed.NewException("type error", "Only bool can be judged.").AddExceptionLine(f.line)
		}

		if !condition.Value() {
			break body
		}

		space, suspend := f.body.ExecWithInit(initSpace, nil)
		defer space.Clear()
		if !nl_interface.IsNil(suspend) {
			switch suspend.InterruptType() {
			case interrupt.BreakInterruptType:
				breaks, yes := interrupt.InterruptFamilyInstance.IsBreak(suspend)
				if !yes {
					return nil, f.seed.NewException("system panic", fmt.Sprintf("BreakInterruptType does not mean a Break anymore.\n%+v", suspend)).AddExceptionLine(f.line)
				}
				if !f.IsMyTag(breaks.Tag()) {
					return nil, suspend
				}
				break body
			case interrupt.ContinueInterruptType:
				continues, yes := interrupt.InterruptFamilyInstance.IsContinue(suspend)
				if !yes {
					return nil, f.seed.NewException("system panic", fmt.Sprintf("ContinueInterruptType does not mean a Continue anymore.\n%+v", suspend)).AddExceptionLine(f.line)
				}
				if !f.IsMyTag(continues.Tag()) {
					return nil, suspend
				}
			default:
				return nil, suspend
			}
		}
		endSpace, suspend := f.end.ExecWithInit(initSpace, nil)
		defer endSpace.Clear()
		if !nl_interface.IsNil(suspend) {
			return nil, suspend
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
		tag.Equal(f.tag) {
		return true
	}
	return false
}

func (f *For) SetCondition(condition concept.Pipe) {
	f.condition = condition
}
func (f *For) Body() concept.CodeBlock {
	return f.body
}
func (f *For) Init() concept.CodeBlock {
	return f.init
}
func (f *For) End() concept.CodeBlock {
	return f.end
}

type ForCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	StringCreator          func(string) concept.String
	BoolCreator            func(bool) concept.Bool
	CodeBlockCreator       func() concept.CodeBlock
	ConstIndexCreator      func(concept.Variable) *index.ConstIndex
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type ForCreator struct {
	Seeds            map[string]func(concept.Pool, *For) (string, concept.Exception)
	param            *ForCreatorParam
	defaultCondition concept.Pipe
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
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *ForCreator) ToLanguage(language string, space concept.Pool, instance *For) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *ForCreator) GetDefaultCondition() concept.Pipe {
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
		Seeds:            map[string]func(concept.Pool, *For) (string, concept.Exception){},
		param:            param,
		defaultCondition: param.ConstIndexCreator(param.BoolCreator(true)),
		defaultTag:       param.StringCreator(""),
	}
}
