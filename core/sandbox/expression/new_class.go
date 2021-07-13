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
	ToLanguage(string, concept.Pool, *NewClass) (string, concept.Exception)
	NewClass() concept.Class
	NewException(string, string) concept.Exception
	NewString(string) concept.String
}

type NewClass struct {
	*adaptor.ExpressionIndex
	items []concept.Pipe
	lines []concept.Line
	seed  NewClassSeed
}

func (f *NewClass) SetItems(items []concept.Pipe) {
	f.items = items
}

func (f *NewClass) SetLines(lines []concept.Line) {
	f.lines = lines
}

func (f *NewClass) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (a *NewClass) ToString(prefix string) string {
	subPrefix := fmt.Sprintf("%v\t", prefix)
	items := []string{}
	for _, line := range a.items {
		items = append(items, fmt.Sprintf("%v%v", subPrefix, line.ToString(subPrefix)))
	}
	return fmt.Sprintf("class {\n%v\n%v}", strings.Join(items, "\n"), prefix)
}

func (a *NewClass) Exec(space concept.Pool) (concept.Variable, concept.Interrupt) {
	class := a.seed.NewClass()
	for cursor, linePre := range a.items {
		lineProvide, yes := index.IndexFamilyInstance.IsProvideIndex(linePre)
		if yes {
			funcsPre, suspend := lineProvide.Get(space)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			funcs, yes := variable.VariableFamilyInstance.IsFunctionHome(funcsPre)
			if !yes {
				return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type in NewClass.provide : %v", funcsPre.Type())).AddExceptionLine(a.lines[cursor])
			}
			class.SetProvide(a.seed.NewString(lineProvide.Name()), funcs)
			continue
		}
		lineRequire, yes := index.IndexFamilyInstance.IsRequireIndex(linePre)
		if yes {
			funcsPre, suspend := lineRequire.Get(space)
			if !nl_interface.IsNil(suspend) {
				return nil, suspend
			}
			funcs, yes := variable.VariableFamilyInstance.IsDefineFunction(funcsPre)
			if !yes {
				return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported variable type in NewClass.require : %v", funcsPre.Type())).AddExceptionLine(a.lines[cursor])
			}
			class.SetRequire(a.seed.NewString(lineRequire.Name()), funcs)
			continue
		}
		return nil, a.seed.NewException("runtime error", fmt.Sprintf("Unsupported index type in NewClass : %v", linePre.Type())).AddExceptionLine(a.lines[cursor])
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
	Seeds map[string]func(concept.Pool, *NewClass) (string, concept.Exception)
	param *NewClassCreatorParam
}

func (s *NewClassCreator) New() *NewClass {
	back := &NewClass{
		items: []concept.Pipe{},
		lines: []concept.Line{},
		seed:  s,
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

func (s *NewClassCreator) ToLanguage(language string, space concept.Pool, instance *NewClass) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func NewNewClassCreator(param *NewClassCreatorParam) *NewClassCreator {
	return &NewClassCreator{
		Seeds: map[string]func(concept.Pool, *NewClass) (string, concept.Exception){},
		param: param,
	}
}
