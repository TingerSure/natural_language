package runtime

import (
	"errors"
	"github.com/TingerSure/natural_language/core/ambiguity"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"os"
)

type Runtime struct {
	lexer           *lexer.Lexer
	grammar         *grammar.Grammar
	ambiguity       *ambiguity.Ambiguity
	libs            *tree.LibraryManager
	languages       *tree.LanguageManager
	box             *sandbox.Sandbox
	rootSpace       *closure.Closure
	defaultLanguage string
}

func (r *Runtime) SetDefaultLanguage(name string) {
	r.defaultLanguage = name

}
func (r *Runtime) GetDefaultLanguage() string {
	return r.defaultLanguage
}

func (r *Runtime) GetLibraryManager() *tree.LibraryManager {
	return r.libs
}
func (r *Runtime) GetLanguageManager() *tree.LanguageManager {
	return r.languages
}

func (r *Runtime) Bind(languageName string) {
	r.languages.GetLanguage(languageName).PackagesIterate(func(_ string, instance tree.Package) bool {
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
		selecteds = append(selecteds, r.ambiguity.Filter(candidates))
	}
	if 0 == len(selecteds) {
		return nil, errors.New("No rules available to match this sentence.")
	}
	return r.ambiguity.Filter(selecteds).Index(), nil
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
		lexer:     lexer.NewLexer(),
		grammar:   grammar.NewGrammar(),
		ambiguity: ambiguity.NewAmbiguity(),
		libs:      tree.NewLibraryManager(),
		languages: tree.NewLanguageManager(),
		rootSpace: closure.NewClosure(nil, &closure.ClosureParam{
			StringCreator: func(value string) concept.String {
				return variable.NewString(value)
			},
			EmptyCreator: func() concept.Null {
				return variable.NewNull()
			},
		}),
	}
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root: runtime.rootSpace,
		OnError: func(err error) {
			param.OnError(err)
		},
		OnPrint: func(value concept.Variable) {
			param.OnPrint(value)
		},
		EventSize: param.EventSize,
	})

	return runtime
}
