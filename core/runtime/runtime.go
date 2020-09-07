package runtime

import (
	"errors"
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/ambiguity"
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
		phrases := river.GetWait().PeekAll()
		phraseString := []string{}
		phraseType := []string{}
		for _, phrase := range phrases {
			phraseString = append(phraseString, phrase.ToString())
			phraseType = append(phraseType, phrase.Types().Name())
		}
		return fmt.Sprintf("%v \nNo struct rule can match ( %v ).", strings.Join(phraseString, ""), strings.Join(phraseType, ", "))
	}

	runtimePriorityErrorFormatDefault = func() string {
		return "No priority rules available to match this sentence."
	}
)

type Runtime struct {
	lexer               *lexer.Lexer
	grammar             *grammar.Grammar
	ambiguity           *ambiguity.Ambiguity
	libs                *LibraryManager
	box                 *sandbox.Sandbox
	rootSpace           *closure.Closure
	defaultLanguage     string
	structErrorFormat   func(*grammar.River) string
	priorityErrorFormat func() string
}

func (r *Runtime) SetStructErrorFormat(format func(*grammar.River) string) {
	if format == nil {
		r.structErrorFormat = runtimeStructErrorFormatDefault
	}
	r.structErrorFormat = format
}

func (r *Runtime) SetPriorityErrorFormat(format func() string) {
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
			r.grammar.AddStructRule(source.GetStructRules())
			r.grammar.AddVocabularyRule(source.GetVocabularyRules())
			r.ambiguity.AddRule(source.GetPriorityRules())
		}
		return false
	})

}

func (r *Runtime) Deal(sentence string) (concept.Index, error) {
	var group *lexer.FlowGroup = r.lexer.Instances(sentence)
	selecteds := []tree.Phrase{}
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

		candidates := []tree.Phrase{}

		valley.Iterate(func(river *grammar.River) bool {
			candidates = append(candidates, river.GetWait().Peek())
			return false
		})

		selecteds = append(selecteds, r.ambiguity.Filter(candidates)...)
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
			if river.GetWait().Len() < min.GetWait().Len() {
				min = river
			}
		}
		return nil, errors.New(r.structErrorFormat(min))
	}
	results := r.ambiguity.Filter(selecteds)
	if 1 != len(results) {
		return nil, errors.New(r.priorityErrorFormat())
	}
	return results[0].Index(), nil
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
		ambiguity:           ambiguity.NewAmbiguity(),
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
