package std

import (
	"github.com/TingerSure/natural_language/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

type StdObject interface {
	Print(concept.Variable)
}

const (
	PrintfContent = "content"
)

var (
	Std StdObject = nil

	Print = variable.NewSystemFunction(
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			if Std != nil || !nl_interface.IsNil(input) {
				Std.Print(input.Get(PrintfContent))
			}
			return input, nil
		},
	)
)
