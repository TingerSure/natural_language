package adaptor

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ExpressionIndexSeed interface {
	Type() string
	NewException(string, string) concept.Exception
}

type ExpressionIndex struct {
	exec func(concept.Closure) (concept.Variable, concept.Interrupt)
	seed ExpressionIndexSeed
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

func (e *ExpressionIndex) Get(space concept.Closure) (concept.Variable, concept.Interrupt) {
	return e.exec(space)
}

func (e *ExpressionIndex) Set(concept.Closure, concept.Variable) concept.Interrupt {
	return e.seed.NewException("read only", "Expression result does not need to be changed.")
}

type ExpressionIndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
}

type ExpressionIndexCreator struct {
	param *ExpressionIndexCreatorParam
}

func (s *ExpressionIndexCreator) New(exec func(concept.Closure) (concept.Variable, concept.Interrupt)) *ExpressionIndex {
	return &ExpressionIndex{
		exec: exec,
		seed: s,
	}
}

func (s *ExpressionIndexCreator) Type() string {
	return IndexExpressionType
}

func (s *ExpressionIndexCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func NewExpressionIndexCreator(param *ExpressionIndexCreatorParam) *ExpressionIndexCreator {
	return &ExpressionIndexCreator{
		param: param,
	}
}
