package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/application/cli"
	"github.com/TingerSure/natural_language/core/ambiguity"
	"github.com/TingerSure/natural_language/core/grammar"
	"github.com/TingerSure/natural_language/core/lexer"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox"
	"github.com/TingerSure/natural_language/core/sandbox/closure"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese"
	"github.com/TingerSure/natural_language/language/chinese/source"
	"github.com/TingerSure/natural_language/library/system"
	"os"
)

func bind(l *lexer.Lexer, g *grammar.Grammar, a *ambiguity.Ambiguity, sources []tree.Source) {
	for _, source := range sources {
		l.AddNaturalSource(source)
		g.AddStructRule(source.GetStructRules())
		g.AddVocabularyRule(source.GetVocabularyRules())
		a.AddRule(source.GetPriorityRules())
	}
}

func getLexerGrammarSandboxAmbiguity(space concept.Closure) (*lexer.Lexer, *grammar.Grammar, *ambiguity.Ambiguity, *sandbox.Sandbox) {

	VM := runtime.NewRuntime()
	VM.GetLibraryManager().AddLibrary("system", system.NewSystemLibrary())

	VM.GetLanguageManager().AddLanguage("chinese", chinese.NewChinese(VM.GetLibraryManager()))

	bind(l, g, a, source.AllRules())
	box := sandbox.NewSandbox(&sandbox.SandboxParam{
		Root: space,
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]: \033[00m%v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
	})
	return l, g, a, box
}

func stringToIndex(info string) concept.Index {
	return index.NewConstIndex(variable.NewString(info))
}

func deal(l *lexer.Lexer, g *grammar.Grammar, a *ambiguity.Ambiguity, sentence string) []concept.Index {
	var group *lexer.FlowGroup = l.Instances(sentence)
	back := []concept.Index{}
	selecteds := []tree.Phrase{}
	for _, flow := range group.GetInstances() {
		// back = append(back, index.NewConstIndex(variable.NewString(flow.ToString())))
		rivers, err := g.Instances(flow)

		if err != nil {
			// back = append(back, index.NewConstIndex(variable.NewString(err.Error())))
			continue
		}
		candidates := []tree.Phrase{}
		for _, river := range rivers {
			// back = append(back, index.NewConstIndex(variable.NewString(river.ToString())))
			candidates = append(candidates, river.GetWait().Peek())
		}
		selecteds = append(selecteds, a.Filter(candidates))
	}
	if 0 == len(selecteds) {
		return append(back, index.NewConstIndex(variable.NewString("No rules available to match this sentence.")))
	}
	selected := a.Filter(selecteds)
	back = append(back, selected.Index())
	return back
}

func test4() {
	rootSpace := closure.NewClosure(nil)
	l, g, a, box := getLexerGrammarSandboxAmbiguity(rootSpace)

	err := box.Start()
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[System]: \033[00m%v\n", err.Error()))
		return
	}
	scan := cli.NewScan(
		os.Stdin,
		func(input string) {
			indexes := deal(l, g, a, input)
			for _, index := range indexes {
				fmt.Printf("\033[1;32m[LOG]: \033[00m%v\n", index.ToString(""))
				box.Exec(index)
			}
		},
		func() {
			os.Stderr.WriteString("\033[1;36m[TS]: \033[00m")
		},
	)
	scan.Run()
}
func main() {
	test4()
}
