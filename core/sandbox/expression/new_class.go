package expression

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/expression/adaptor"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

type NewClassSeed interface {
	ToLanguage(string, *NewClass) string
	NewClass() concept.Class
	NewException(string, string) concept.Exception
	NewString(string) concept.String
}

type NewClass struct {
	*adaptor.ExpressionIndex
	lines []concept.Index
	seed  NewClassSeed
}

func (f *NewClass) SetLines(lines []concept.Index) {
	f.lines = lines
}

func (f *NewClass) ToLanguage(language string) string {
	return f.seed.ToLanguage(language, f)
}

func (a *NewClass) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	lines := []string{}
	for _, line := range a.lines {
		lines = append(lines, fmt.Sprintf("%v%v", subPrefix, line.ToString(subPrefix)))
	}
	return fmt.Sprintf("class {\n%v\n%v}", strings.Join(lines, "\n"), prefix)
}

func (a *NewClass) Anticipate(space concept.Closure) concept.Variable {
	class := a.seed.NewClass()
	for _, linePre := range a.lines {
		lineProvide, yes := index.IndexFamilyInstance.IsProvideIndex(linePre)
		if yes {
			funcsPre := lineProvide.Anticipate(space)
			funcs, yes := variable.VariableFamilyInstance.IsFunctionHome(funcsPre)
			if yes {
				class.SetProvide(a.seed.NewString(lineProvide.Name()), funcs)
			}
			continue
		}
		lineRequire, yes := index.IndexFamilyInstance.IsRequireIndex(linePre)
		if yes {
			funcsPre := lineRequire.Anticipate(space)
			funcs, yes := variable.VariableFamilyInstance.IsDefineFunction(funcsPre)
			if yes {
				class.SetRequire(a.seed.NewString(lineRequire.Name()), funcs)
			}
			continue
		}
	}
	return class
}

func (a *NewClass) Exec(space concept.Closure) (concept.Variable, concept.Interrupt) {
	class := a.seed.NewClass()
	for _, linePre := range a.lines {
		lineProvide, yes := index.IndexFamilyInstance.IsProvideIndex(linePre)
		if yes {
			funcsPre, suspend := lineProvide.Get(space)
			if nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			funcs, yes := variable.VariableFamilyInstance.IsFunctionHome(funcsPre)
			if !yes {
				return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type in NewClass.provide : %v", funcsPre.Type()))
			}
			class.SetProvide(a.seed.NewString(lineProvide.Name()), funcs)
			continue
		}
		lineRequire, yes := index.IndexFamilyInstance.IsRequireIndex(linePre)
		if yes {
			funcsPre, suspend := lineRequire.Get(space)
			if nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			funcs, yes := variable.VariableFamilyInstance.IsDefineFunction(funcsPre)
			if !yes {
				return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type in NewClass.require : %v", funcsPre.Type()))
			}
			class.SetRequire(a.seed.NewString(lineRequire.Name()), funcs)
			continue
		}
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported index type in NewClass : %v", linePre.Type()))
	}
	return class, nil
}

type NewClassCreatorParam struct {
	ExpressionIndexCreator func(concept.Expression) *adaptor.ExpressionIndex
	ExceptionCreator       func(string, string) concept.Exception
	StringCreator          func(string) concept.String
	ClassCreator           func() concept.Class
}

type NewClassCreator struct {
	Seeds map[string]func(string, *NewClass) string
	param *NewClassCreatorParam
}

func (s *NewClassCreator) New() *NewClass {
	back := &NewClass{
		seed: s,
	}
	back.ExpressionIndex = s.param.ExpressionIndexCreator(back)
	return back
}

func (s *NewClassCreator) NewClass() concept.Class {
	return s.param.ClassCreator()
}

func (s *NewClassCreator) NewException(name string, message string) concept.Exception {
	return s.param.ExceptionCreator(name, message)
}

func (s *NewClassCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func (s *NewClassCreator) ToLanguage(language string, instance *NewClass) string {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString("")
	}
	return seed(language, instance)
}

func NewNewClassCreator(param *NewClassCreatorParam) *NewClassCreator {
	return &NewClassCreator{
		Seeds: map[string]func(string, *NewClass) string{},
		param: param,
	}
}
