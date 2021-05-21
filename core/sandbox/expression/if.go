package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/code_block"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type IfSeed interface {
	ToLanguage(string, *If) string
	NewException(string, string) concept.Exception
	NewNull() concept.Null
	NewClosure(concept.Closure) concept.Closure
}

type If struct {
	*adaptor.ExpressionIndex
	condition concept.Index
	primary   *code_block.CodeBlock
	secondary *code_block.CodeBlock
	seed      IfSeed
}

func (f *If) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (f *If) ToString(prefix string) string {
	primaryToString := fmt.Sprintf("if (%v) %v", f.condition.ToString(prefix), f.primary.ToString(prefix))
	if f.secondary.Size() == 0 {
		return primaryToString
	}
	return fmt.Sprintf("%v else %v", primaryToString, f.secondary.ToString(prefix))
}
func (e *If) Anticipate(space concept.Closure) concept.Variable {
	return e.seed.NewNull()
}
func (f *If) Exec(parent concept.Closure) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		return nil, f.seed.NewException("system error", "No condition for judgment.")
	}
	initSpace := f.seed.NewClosure(parent)
	defer initSpace.Clear()

	preCondition, suspend := f.condition.Get(initSpace)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return nil, f.seed.NewException("type error", "Only bool can be judged.")
	}

	var active *code_block.CodeBlock
	if condition.Value() {
		active = f.primary
	} else {
		active = f.secondary
	}

	space, suspend := active.Exec(initSpace, nil)
	defer space.Clear()
	return f.seed.NewNull(), suspend
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

type IfCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	CodeBlockCreator       func() *code_block.CodeBlock
	ClosureCreator         func(concept.Closure) concept.Closure
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type IfCreator struct {
	Seeds            map[string]func(string, *If) string
	param            *IfCreatorParam
	defaultCondition concept.Index
	defaultTag       concept.String
}

func (s *IfCreator) New() *If {
	back := &If{
		primary:   s.param.CodeBlockCreator(),
		secondary: s.param.CodeBlockCreator(),
		seed:      s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *IfCreator) ToLanguage(language string, instance *If) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func (s *IfCreator) NewClosure(parent concept.Closure) concept.Closure {
	return s.param.ClosureCreator(parent)
}

func (s *IfCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *IfCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewIfCreator(param *IfCreatorParam) *IfCreator {
	return &IfCreator{
		Seeds: map[string]func(string, *If) string{},
		param: param,
	}
}
