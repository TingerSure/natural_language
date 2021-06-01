package runtime

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func NewPoolObject(libs *tree.LibraryManager, pool concept.Pool) concept.Object {
	object := libs.Sandbox.Variable.Object.New()
	keyParam := libs.Sandbox.Variable.String.New("key")
	valueParam := libs.Sandbox.Variable.String.New("value")
	object.SetField(
		libs.Sandbox.Variable.String.New("GetLocal"),
		libs.Sandbox.Variable.SystemFunction.New(
			func(input concept.Param, _ concept.Variable) (concept.Param, concept.Exception) {
				keyPre := input.Get(keyParam)
				key, yes := variable.VariableFamilyInstance.IsStringHome(keyPre)
				if !yes {
					return nil, libs.Sandbox.Variable.Exception.NewOriginal("type error", fmt.Sprintf("Key is not a string: %v", keyPre.ToString("")))
				}
				back, suspend := pool.GetLocal(key)
				if !nl_interface.IsNil(suspend) {
					return nil, suspend
				}
				output := libs.Sandbox.Variable.Param.New()
				output.Set(valueParam, back)
				return output, nil
			},
			nil,
			[]concept.String{keyParam},
			[]concept.String{valueParam},
		),
	)
	return object
}
