package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	// "github.com/TingerSure/natural_language/language/chinese"
	// "github.com/TingerSure/natural_language/library/system"
	// "github.com/TingerSure/natural_language/library/system/std"
	"os"
	"time"
)

const (
	ChineseName = "chinese"
)

func getVM() (*runtime.Runtime, error) {
	VM := runtime.NewRuntime(&runtime.RuntimeParam{
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]:\033[00m %v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
		Roots: []string{
			"./",
		},
	})

	// system.BindSystem(VM.GetLibraryManager(), &system.SystemLibraryParam{
	// 	Std: &std.StdParam{
	// 		Error: func(value concept.Variable) {
	// 			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]:\033[00m %v\n", value.ToLanguage(ChineseName)))
	// 		},
	// 		Print: func(value concept.Variable) {
	// 			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToLanguage(ChineseName)))
	// 		},
	// 	},
	// })
	// chinese.BindRule(VM.GetLibraryManager(), ChineseName)
	// chinese.BindLanguage(VM.GetLibraryManager(), ChineseName)
	err := VM.Read("test2.nl")
	if err != nil {
		return nil, err
	}
	VM.Bind()
	VM.SetDefaultLanguage(ChineseName)
	return VM, nil
}

func test() {

	VM, err := getVM()
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[System]:\033[00m %v\n", err.Error()))
		return
	}
	err = VM.Start()
	if err != nil {
		os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[System]:\033[00m %v\n", err.Error()))
		return
	}
	cli := runtime.NewScan(&runtime.ScanParam{
		Stream: os.Stdin,
		OnReader: func(input string) bool {
			startTime := time.Now().Unix()
			index, err := VM.Deal(input)

			fmt.Printf("\033[1;32m[TIME]:\033[00m %vs\n", time.Now().Unix()-startTime)

			if err != nil {
				fmt.Printf("\033[1;32m[ERROR]:\033[00m %v\n", err.Error())
				return true
			}

			fmt.Printf("\033[1;32m[LOG]:\033[00m %v\n", index.ToString(""))
			VM.Exec(index)
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
