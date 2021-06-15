package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type IfSeed interface {
	ToLanguage(string, concept.Pool, *If) (string, concept.Exception)
	NewException(string, string) concept.Exception
	NewNull() concept.Null
	NewPool(concept.Pool) concept.Pool
}

type If struct {
	*adaptor.ExpressionIndex
	condition concept.Pipe
	primary   concept.CodeBlock
	secondary concept.CodeBlock
	line      concept.Line
	seed      IfSeed
}

func (f *If) SetLine(line concept.Line) {
	f.line = line
}

func (f *If) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (f *If) ToString(prefix string) string {
	primaryToString := fmt.Sprintf("if (%v) %v", f.condition.ToString(prefix), f.primary.ToString(prefix))
	if f.secondary.Size() == 0 {
		return primaryToString
	}
	return fmt.Sprintf("%v else %v", primaryToString, f.secondary.ToString(prefix))
}
func (e *If) Anticipate(space concept.Pool) concept.Variable {
	return e.seed.NewNull()
}
func (f *If) Exec(parent concept.Pool) (concept.Variable, concept.Interrupt) {

	if nl_interface.IsNil(f.condition) {
		return nil, f.seed.NewException("system error", "No condition for judgment.").AddLine(f.line)
	}
	initSpace := f.seed.NewPool(parent)
	defer initSpace.Clear()

	preCondition, suspend := f.condition.Get(initSpace)
	if !nl_interface.IsNil(suspend) {
		return nil, suspend
	}

	condition, yes := variable.VariableFamilyInstance.IsBool(preCondition)
	if !yes {
		return nil, f.seed.NewException("type error", "Only bool can be judged.").AddLine(f.line)
	}

	var active concept.CodeBlock
	if condition.Value() {
		active = f.primary
	} else {
		active = f.secondary
	}

	space, suspend := active.ExecWithInit(initSpace, nil)
	defer space.Clear()
	return f.seed.NewNull(), suspend
}

func (f *If) SetCondition(condition concept.Pipe) {
	f.condition = condition
}

func (f *If) Primary() concept.CodeBlock {
	return f.primary
}

func (f *If) Secondary() concept.CodeBlock {
	return f.secondary
}

type IfCreatorParam struct {
	ExceptionCreator       func(string, string) concept.Exception
	CodeBlockCreator       func() concept.CodeBlock
	PoolCreator            func(concept.Pool) concept.Pool
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	NullCreator            func() concept.Null
}

type IfCreator struct {
	Seeds            map[string]func(concept.Pool, *If) (string, concept.Exception)
	param            *IfCreatorParam
	defaultCondition concept.Pipe
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

func (s *IfCreator) ToLanguage(language string, space concept.Pool, instance *If) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *IfCreator) NewPool(parent concept.Pool) concept.Pool {
	return s.param.PoolCreator(parent)
}

func (s *IfCreator) NewNull() concept.Null {
	return s.param.NullCreator()
}

func (s *IfCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewIfCreator(param *IfCreatorParam) *IfCreator {
	return &IfCreator{
		Seeds: map[string]func(concept.Pool, *If) (string, concept.Exception){},
		param: param,
	}
}
