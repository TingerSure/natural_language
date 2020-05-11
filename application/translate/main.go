package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/language/chinese"
	"github.com/TingerSure/natural_language/language/english"
	"github.com/TingerSure/natural_language/library/system"
	"github.com/TingerSure/natural_language/library/system/std"
	"os"
)

const (
	ChineseName = "chinese"
	EnglishName = "english"
)

func getVM() *runtime.Runtime {
	VM := runtime.NewRuntime(&runtime.RuntimeParam{
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]:\033[00m %v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
	})
	VM.GetLibraryManager().AddSystemLibrary(system.NewSystemLibrary(VM.GetLibraryManager(), &system.SystemLibraryParam{
		Std: &std.StdParam{
			Error: func(value concept.Variable) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]:\033[00m %v\n", value.ToString("")))
			},
			Print: func(value concept.Variable) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
			},
		},
	}))
	VM.GetLibraryManager().AddLibrary(ChineseName, chinese.NewChinese(VM.GetLibraryManager(), ChineseName))
	chinese.ChineseBindLanguage(VM.GetLibraryManager(), ChineseName)
	english.EnglishBindLanguage(VM.GetLibraryManager(), EnglishName)
	VM.Bind()
	VM.SetDefaultLanguage(ChineseName)
	return VM
}

func test() {

	VM := getVM()
	err := VM.Start()
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[System]:\033[00m %v\n", err.Error()))
		return
	}
	cli := runtime.NewScan(&runtime.ScanParam{
		Stream: os.Stdin,
		OnReader: func(input string) bool {
			index, err := VM.Deal(input)
			if err != nil {
				fmt.Printf("\033[1;32m[ERROR]:\033[00m %v\n", err.Error())
				return true
			}

			fmt.Printf("\033[1;32m[TRANSLATE]:\033[00m %v\n", index.ToLanguage(ChineseName))
			fmt.Printf("\033[1;32m[TRANSLATE]:\033[00m %v\n", index.ToLanguage(EnglishName))

			return true
		},
		BeforeReader: func() {
			os.Stderr.WriteString("\033[1;36m[TS]:\033[00m ")
		},
	})
	cli.Run()
}

func main() {
	test()
}
