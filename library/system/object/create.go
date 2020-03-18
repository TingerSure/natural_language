package object

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

var (
	CreateContent = variable.NewString("object")

	Create *variable.SystemFunction = nil
)

func init() {
	Create = variable.NewSystemFunction(
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return variable.NewParam().Set(CreateContent, variable.NewObject()), nil
		},
		[]concept.String{},
		[]concept.String{
			CreateContent,
		},
	)
}
