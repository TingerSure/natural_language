package code_block

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"strings"
)

type CodeBlockSeed interface {
	NewPool(concept.Pool) concept.Pool
}

type CodeBlock struct {
	flow []concept.Pipe
	seed CodeBlockSeed
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
func (c *CodeBlock) AddStep(steps ...concept.Pipe) {
	c.flow = append(c.flow, steps...)
}

func (f *CodeBlock) Exec(
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
		value, suspend := step.Get(space)
		if !nl_interface.IsNil(suspend) {
			return space, suspend
		}
		space.AddExtempore(step, value)
	}
	return space, nil
}

type CodeBlockCreatorParam struct {
	PoolCreator func(concept.Pool) concept.Pool
}

type CodeBlockCreator struct {
	Seeds map[string]func(string, concept.Pool, *CodeBlock) string
	param *CodeBlockCreatorParam
}

func (s *CodeBlockCreator) New() *CodeBlock {
	return &CodeBlock{
		seed: s,
	}
}

func (s *CodeBlockCreator) ToLanguage(language string, space concept.Pool, instance *CodeBlock) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, space, instance)
}

func (s *CodeBlockCreator) NewPool(parent concept.Pool) concept.Pool {
	return s.param.PoolCreator(parent)
}

func NewCodeBlockCreator(param *CodeBlockCreatorParam) *CodeBlockCreator {
	return &CodeBlockCreator{
		Seeds: map[string]func(string, concept.Pool, *CodeBlock) string{},
		param: param,
	}
}
