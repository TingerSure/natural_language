package runtime

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
	"strings"
)

var (
	runtimeStructErrorFormatDefault = func(river *grammar.River) string {
		phrases := river.GetLake().PeekAll()
		phraseString := []string{}
		phraseType := []string{}
		for _, phrase := range phrases {
			phraseString = append(phraseString, phrase.ToString())
			phraseType = append(phraseType, phrase.Types().Name())
		}
		return fmt.Sprintf("%v \nNo struct rule can match ( %v ).", strings.Join(phraseString, ""), strings.Join(phraseType, ", "))
	}

	runtimePriorityErrorFormatDefault = func(rivers []*grammar.River) string {
		riverString := []string{}
		for _, river := range rivers {
			riverString = append(riverString, river.GetLake().Peek().ToString())
		}
		return fmt.Sprintf("%v \nNo priority rule can distinguish the above meanings.", strings.Join(riverString, "\n"))
	}
)

type Runtime struct {
	lexer               *lexer.Lexer
	grammar             *grammar.Grammar
	libs                *LibraryManager
	box                 *sandbox.Sandbox
	rootSpace           *closure.Closure
	defaultLanguage     string
	structErrorFormat   func(*grammar.River) string
	priorityErrorFormat func([]*grammar.River) string
}

func (r *Runtime) SetStructErrorFormat(format func(*grammar.River) string) {
	if format == nil {
		r.structErrorFormat = runtimeStructErrorFormatDefault
	}
	r.structErrorFormat = format
}

func (r *Runtime) SetPriorityErrorFormat(format func([]*grammar.River) string) {
	if format == nil {
		r.priorityErrorFormat = runtimePriorityErrorFormatDefault
	}
	r.priorityErrorFormat = format
}

func (r *Runtime) SetDefaultLanguage(name string) {
	r.defaultLanguage = name
}

func (r *Runtime) GetDefaultLanguage() string {
	return r.defaultLanguage
}

func (r *Runtime) GetLibraryManager() *LibraryManager {
	return r.libs
}

func (r *Runtime) Bind() {
	r.libs.PageIterate(func(instance tree.Page) bool {
		for _, source := range instance.GetSources() {
			r.lexer.AddNaturalSource(source)
			r.grammar.GetReach().AddRule(source.GetStructRules())
			r.grammar.GetSection().AddRule(source.GetVocabularyRules())
			r.grammar.GetDam().AddRule(source.GetPriorityRules())
		}
		return false
	})
}

func (r *Runtime) Deal(sentence string) (concept.Index, error) {
	var group *lexer.FlowGroup = r.lexer.Instances(sentence)
	selecteds := []*grammar.River{}
	mostMatch := []*grammar.River{}
	for _, flow := range group.GetInstances() {
		sourceValley, err := r.grammar.Instances(flow)
		if err != nil {
			continue
		}

		valley, min, err := sourceValley.Filter()
		if err != nil {
			if !nl_interface.IsNil(min) {
				mostMatch = append(mostMatch, min)
			}
			continue
		}
		selecteds = append(selecteds, valley.AllRivers()...)
	}
	if 0 == len(selecteds) {

		if 0 == len(mostMatch) {
			return nil, errors.New("Empty sentence.")
		}

		var min *grammar.River
		for _, river := range mostMatch {
			if nl_interface.IsNil(min) {
				min = river
				continue
			}
			if river.GetLake().Len() < min.GetLake().Len() {
				min = river
			}
		}
		return nil, errors.New(r.structErrorFormat(min))
	}
	results := r.grammar.GetDam().Filter(selecteds)
	if 1 != len(results) {
		return nil, errors.New(r.priorityErrorFormat(results))
	}

	return results[0].GetLake().Peek().Index(), nil
}

func (r *Runtime) Read(stream *os.File) error {
	var scanErr error = nil
	NewScan(&ScanParam{
		Stream: stream,
		OnReader: func(input string) bool {
			index, err := r.Deal(input)
			if err != nil {
				scanErr = err
				return false
			}
			r.Exec(index)
			return true
		},
		BeforeReader: func() {},
	}).Run()

	return scanErr
}

func (r *Runtime) Start() error {
	return r.box.Start()
}

func (r *Runtime) Stop() error {
	return r.box.Stop()
}

func (r *Runtime) Exec(hand concept.Index) {
	r.box.Exec(hand)
}

type RuntimeParam struct {
	OnError   func(error)
	OnPrint   func(concept.Variable)
	EventSize int
}

func NewRuntime(param *RuntimeParam) *Runtime {
	runtime := &Runtime{
		lexer:               lexer.NewLexer(),
		grammar:             grammar.NewGrammar(),
		libs:                NewLibraryManager(),
		structErrorFormat:   runtimeStructErrorFormatDefault,
		priorityErrorFormat: runtimePriorityErrorFormatDefault,
	}
	runtime.rootSpace = runtime.libs.Sandbox.Closure.New(nil)
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root:      runtime.rootSpace,
		OnError:   param.OnError,
		OnPrint:   param.OnPrint,
		EventSize: param.EventSize,
	})

	return runtime
}
