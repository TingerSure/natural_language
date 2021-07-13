package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newFindPipeCache(libs *tree.LibraryManager, rootPipeCache *tree.PipeCache) concept.Function {
	matchParam := libs.Sandbox.Variable.String.New("match")
	line := tree.NewLine("[find_pipe_cache]", "")
	return libs.Sandbox.Variable.SystemFunction.New(
		func(param concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
			matchPre := param.Get(matchParam)
			match, yes := variable.VariableFamilyInstance.IsFunctionHome(matchPre)
			if !yes {
				return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Param match is not a function: %v", matchPre.ToString("")))
			}
			var exception concept.Exception = nil
			rootPipeCache.Iterate(func(pipe concept.Function, value concept.Variable) bool {
				input := libs.Sandbox.Variable.Param.New()
				input.SetOriginal("pipe", pipe)
				input.SetOriginal("value", value)
				output, exception := match.Exec(input, nil)
				if !nl_interface.IsNil(exception) {
					return true
				}
				stopPre := output.GetOriginal("stop")
				if stopPre.IsNull() {
					return false
				}
				stop, yes := variable.VariableFamilyInstance.IsBool(stopPre)
				if !yes {
					exception = libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Return stop is not a bool: %v", stopPre.ToString(""))).AddExceptionLine(line)
					return true
				}
				return stop.Value()
			})
			if !nl_interface.IsNil(exception) {
				return nil, exception.AddExceptionLine(line)
			}
			return libs.Sandbox.Variable.Param.New(), nil
		},
		[]concept.String{matchParam},
		[]concept.String{},
	)
}
