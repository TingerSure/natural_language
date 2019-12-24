package object

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

const (
	CreateContent = "object"
)

var (
	Create *variable.SystemFunction = nil
)

func init() {
	Create = variable.NewSystemFunction(
		func(_ concept.Param, _ concept.Object) (concept.Param, concept.Exception) {
			return variable.NewParamWithInit(map[string]concept.Variable{
				CreateContent: variable.NewObject(),
			}), nil
		},
		[]string{},
		[]string{
			CreateContent,
		},
	)
}
