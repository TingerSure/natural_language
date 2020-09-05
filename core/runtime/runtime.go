package runtime

import (
	"errors"
	"github.com/TingerSure/natural_language/core/ambiguity"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
)

var (
	runtimeStructErrorFormatDefault = func() string {
		return "No struct rules available to match this sentence."
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
	structErrorFormat   func() string
	priorityErrorFormat func() string
}

func (r *Runtime) SetStructErrorFormat(format func() string) {
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
	for _, flow := range group.GetInstances() {
		rivers, err := r.grammar.Instances(flow)

		if err != nil {
			continue
		}
		candidates := []tree.Phrase{}
		for _, river := range rivers {
			candidates = append(candidates, river.GetWait().Peek())
		}
		selecteds = append(selecteds, r.ambiguity.Filter(candidates)...)
	}
	if 0 == len(selecteds) {
		return nil, errors.New(r.structErrorFormat())
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
