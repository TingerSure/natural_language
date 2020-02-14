package runtime

import (
	"github.com/TingerSure/natural_language/core/ambiguity"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type Runtime struct {
	lexer     *lexer.Lexer
	grammar   *grammar.Grammar
	ambiguity *ambiguity.Ambiguity
	libs      *tree.LibraryManager
	languages *tree.LanguageManager
	box       *sandbox.Sandbox
	rootSpace *closure.Closure
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

func (r *Runtime) Deal(sentence string) []concept.Index {
	var group *lexer.FlowGroup = r.lexer.Instances(sentence)
	back := []concept.Index{}
	selecteds := []tree.Phrase{}
	for _, flow := range group.GetInstances() {
		// back = append(back, index.NewConstIndex(variable.NewString(flow.ToString())))
		rivers, err := r.grammar.Instances(flow)

		if err != nil {
			// back = append(back, index.NewConstIndex(variable.NewString(err.Error())))
			continue
		}
		candidates := []tree.Phrase{}
		for _, river := range rivers {
			// back = append(back, index.NewConstIndex(variable.NewString(river.ToString())))
			candidates = append(candidates, river.GetWait().Peek())
		}
		selecteds = append(selecteds, r.ambiguity.Filter(candidates))
	}
	if 0 == len(selecteds) {
		return append(back, index.NewConstIndex(variable.NewString("No rules available to match this sentence.")))
	}
	selected := r.ambiguity.Filter(selecteds)
	back = append(back, selected.Index())
	return back
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
		rootSpace: closure.NewClosure(nil),
	}
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root: runtime.rootSpace,
		OnError: func(err error) {
			param.OnError(err)
			// os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]: \033[00m%v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			param.OnPrint(value)

			// os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: param.EventSize,
	})

	return runtime
}
