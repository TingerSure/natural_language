package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/application/cli"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/language/chinese"
	"github.com/TingerSure/natural_language/library/system"
	"github.com/TingerSure/natural_language/library/system/std"
	"os"
)

func getVM() *runtime.Runtime {
	VM := runtime.NewRuntime(&runtime.RuntimeParam{
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]: \033[00m%v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
	})
	VM.GetLibraryManager().AddSystemLibrary(system.NewSystemLibrary(&system.SystemLibraryParam{
		Std: &std.StdParam{
			Error: func(value concept.Variable) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]: \033[00m%v\n", value.ToString("")))
			},
			Print: func(value concept.Variable) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
			},
		},
	}))
	VM.GetLanguageManager().AddLanguage("chinese", chinese.NewChinese(VM.GetLibraryManager()))
	VM.Bind("chinese")
	return VM
}

func test4() {

	VM := getVM()
	err := VM.Start()
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[System]: \033[00m%v\n", err.Error()))
		return
	}
	scan := cli.NewScan(
		os.Stdin,
		func(input string) {
			indexes := VM.Deal(input)
			for _, index := range indexes {
				fmt.Printf("\033[1;32m[LOG]: \033[00m%v\n", index.ToString(""))
				VM.Exec(index)
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
