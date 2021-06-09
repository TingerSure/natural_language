package bind

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func newValueLanguage(libs *tree.LibraryManager) (valueLanguage *variable.ValueLanguageFunction) {
	valueParam := libs.Sandbox.Variable.String.New("value")
	valueLanguage = libs.Sandbox.Variable.ValueLanguageFunction.New(
		[]concept.String{
			valueParam,
		},
		[]concept.String{},
	)
	valueLanguage.SetLanguageOnCallDefaultSeed(func(language string, funcs concept.Function, pool concept.Pool, name string, params concept.Param) (string, concept.Exception) {
		return params.Get(valueParam).(concept.String).Value(), nil
	})
	return
}
