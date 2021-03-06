package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system"
	lib_runtime "github.com/TingerSure/natural_language/library/system/runtime"
	"github.com/TingerSure/natural_language/library/system/std"
	"os"
)

const (
	ChineseName = "chinese"
)

func getVM() (*runtime.Runtime, error) {
	VMParam := &runtime.RuntimeParam{
		OnError: func(err error) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[NL]:\033[00m %v\n", err.Error()))
		},
		OnPrint: func(value concept.Variable) {
			os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", value.ToString("")))
		},
		EventSize: 1024,
		SourceRoots: []string{
			"./",
		},
		SourceExtension: ".nl",
	}
	VM := runtime.NewRuntime(VMParam)

	system.BindSystem(VM.GetLibraryManager(), &system.SystemLibraryParam{
		StdParam: &std.StdParam{
			Error: func(value concept.Variable) {
				message, suspend := value.ToLanguage(ChineseName, nil)
				if !nl_interface.IsNil(suspend) {
					VMParam.OnError(suspend)
					return
				}
				os.Stdout.WriteString(fmt.Sprintf("\033[1;31m[ERR]:\033[00m %v\n", message))
			},
			Print: func(value concept.Variable) {
				message, suspend := value.ToLanguage(ChineseName, nil)
				if !nl_interface.IsNil(suspend) {
					VMParam.OnError(suspend)
					return
				}
				os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[NL]:\033[00m %v\n", message))
			},
			Log: func(value concept.Variable) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;35m[LOG]:\033[00m %v\n", value.ToString("")))
			},
		},
		RuntimeParam: &lib_runtime.RuntimeParam{
			RootSpace:     VM.GetRootSpace(),
			RootPipeCache: VM.GetRootPipeCache(),
		},
	})
	err := VM.Read("main/main")
	if err != nil {
		return nil, err
	}
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
			index, err := VM.Deal(input)
			if err != nil {
				fmt.Printf("\033[1;32m[ERROR]:\033[00m %v\n", err.Error())
				return true
			}
			message, suspend := index.ToCallLanguage(ChineseName, VM.GetRootSpace(), "", nil)
			if !nl_interface.IsNil(suspend) {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[AUTO-LOG-ERROR]:\033[00m %v\n", suspend.Error()))
			} else {
				os.Stdout.WriteString(fmt.Sprintf("\033[1;36m[AUTO-LOG]:\033[00m %v\n", message))
			}
			VM.Exec(index, tree.NewLine(fmt.Sprintf("[scan_input]:%v", input), ""))
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
