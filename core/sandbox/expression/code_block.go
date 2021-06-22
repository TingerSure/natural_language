package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"strings"
)

type CodeBlockSeed interface {
	ToLanguage(string, concept.Pool, *CodeBlock) (string, concept.Exception)
	NewPool(concept.Pool) concept.Pool
}

type CodeBlock struct {
	*adaptor.ExpressionIndex
	flow []concept.Pipe
	seed CodeBlockSeed
}

func (f *CodeBlock) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (c *CodeBlock) Size() int {
	return len(c.flow)
}

const (
	CodeBlockFlowNoneSize     = 0
	CodeBlockFlowOnlyOneSize  = 1
	CodeBlockFlowOnlyOneIndex = 0
)

func (c *CodeBlock) ToStringSimplify(prefix string) string {
	if c.Size() == CodeBlockFlowNoneSize {
		return ""
	}
	if c.Size() == CodeBlockFlowOnlyOneSize {
		return c.flow[CodeBlockFlowOnlyOneIndex].ToString(prefix)
	}
	return c.ToString(prefix)
}

func (c *CodeBlock) ToString(prefix string) string {
	flowPrefix := fmt.Sprintf("%v\t", prefix)
	flowToStrings := make([]string, 0, c.Size())
	for _, flow := range c.flow {
		flowToStrings = append(flowToStrings, fmt.Sprintf("%v%v;", flowPrefix, flow.ToString(flowPrefix)))
	}
	return fmt.Sprintf("{\n%v\n%v}", strings.Join(flowToStrings, "\n"), prefix)
}

func (c *CodeBlock) Steps() []concept.Pipe {
	return c.flow
}

func (c *CodeBlock) AddStep(steps ...concept.Pipe) {
	c.flow = append(c.flow, steps...)
}

func (a *CodeBlock) Anticipate(space concept.Pool) concept.Variable {
	return a.seed.NewPool(space)
}

func (a *CodeBlock) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	return a.ExecWithInit(space, nil)
}

func (f *CodeBlock) ExecWithInit(
	parent concept.Pool,
	init func(concept.Pool) concept.Interrupt,
) (concept.Pool, concept.Interrupt) {
	space := f.seed.NewPool(parent)
	if init != nil {
		suspend := init(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
	}
	for _, step := range f.flow {
		_, suspend := step.Get(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
	}
	return space, nil
}

type CodeBlockCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	PoolCreator            func(concept.Pool) concept.Pool
}

type CodeBlockCreator struct {
	Seeds map[string]func(concept.Pool, *CodeBlock) (string, concept.Exception)
	param *CodeBlockCreatorParam
}

func (s *CodeBlockCreator) New() *CodeBlock {
	back := &CodeBlock{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *CodeBlockCreator) ToLanguage(language string, space concept.Pool, instance *CodeBlock) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *CodeBlockCreator) NewPool(parent concept.Pool) concept.Pool {
	return s.param.PoolCreator(parent)
}

func NewCodeBlockCreator(param *CodeBlockCreatorParam) *CodeBlockCreator {
	return &CodeBlockCreator{
		Seeds: map[string]func(concept.Pool, *CodeBlock) (string, concept.Exception){},
		param: param,
	}
}
