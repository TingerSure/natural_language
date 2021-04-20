package adaptor

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ExpressionIndexSeed interface {
	Type() string
	NewException(string, string) concept.Exception
	NewParam() concept.Param
}

type ExpressionIndex struct {
	expression concept.Expression
	seed       ExpressionIndexSeed
}

var (
	IndexExpressionType = "Expression"
)

func (e *ExpressionIndex) Type() string {
	return e.seed.Type()
}

func (e *ExpressionIndex) SubCodeBlockIterate(func(concept.Index) bool) bool {
	return false
}

func (s *ExpressionIndex) Call(space concept.Closure, param concept.Param) (concept.Param, concept.Exception) {
	funcs, interrupt := s.Get(space)
	if !nl_interface.IsNil(interrupt) {
		return nil, interrupt.(concept.Exception)
	}
	if !funcs.IsFunction() {
		return nil, s.seed.NewException("runtime error", "The result of this expression is not a function.")
	}
	return funcs.(concept.Function).Exec(param, nil)
}

func (s *ExpressionIndex) CallAnticipate(space concept.Closure, param concept.Param) concept.Param {
	funcs := s.expression.Anticipate(space)
	if !funcs.IsFunction() {
		return s.seed.NewParam()
	}
	return funcs.(concept.Function).Anticipate(param, nil)
}

func (e *ExpressionIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return e.expression.Exec(space)
}

func (e *ExpressionIndex) Set(concept.Closure, concept.Variable) concept.Interrupt {
	return e.seed.NewException("read only", "Expression result does not need to be changed.")
}

type ExpressionIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
}

type ExpressionIndexCreator struct {
	param *ExpressionIndexCreatorParam
}

func (s *ExpressionIndexCreator) New(expression concept.Expression) *ExpressionIndex {
	return &ExpressionIndex{
		expression: expression,
		seed:       s,
	}
}

func (s *ExpressionIndexCreator) Type() string {
	return IndexExpressionType
}

func (s *ExpressionIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *ExpressionIndexCreator) NewParam() concept.Param {
	return s.param.ParamCreator()
}

func NewExpressionIndexCreator(param *ExpressionIndexCreatorParam) *ExpressionIndexCreator {
	return &ExpressionIndexCreator{
		param: param,
	}
}
