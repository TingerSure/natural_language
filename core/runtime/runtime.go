package runtime

import (
	"github.com/TingerSure/natural_language/core/ambiguity"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/sanbox"
	"github.com/TingerSure/natural_language/core/sanbox/closure"
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

func NewRuntime() Runtime {

	runtime := &Runtime{
		lexer:     lexer.NewLexer(),
		grammar:   grammar.NewGrammar(),
		ambiguity: ambiguity.NewAmbiguity(),
		libs:      tree.NewLibraryManager(),
		languages: tree.NewLanguageManager(),
		rootSpace: closure.NewClosure(nil),
	}
	runtime.box = sandbox.NewSandbox(&sandbox.SandboxParam{
		Root: space,
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]: \033[00m%v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
	})

	return runtime
}
