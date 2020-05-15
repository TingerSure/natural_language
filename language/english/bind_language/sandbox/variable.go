package sandbox

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"strings"
)

func VariableBindLanguage(libs *runtime.LibraryManager, language string) {
	variable.StringLanguageSeeds[language] = func(language string, instance *variable.String) string {
		return instance.GetLanguage(language)
	}

	variable.NumberLanguageSeeds[language] = func(language string, instance *variable.Number) string {
		return fmt.Sprintf("%v", instance.Value())
	}

	variable.FunctionLanguageSeeds[language] = func(language string, instance *variable.Function) string {
		return instance.Name().ToLanguage(language)
	}

	variable.ParamLanguageSeeds[language] = func(language string, instance *variable.Param) string {
		items := []string{}

		instance.Iterate(func(key concept.String, value concept.Variable) bool {
			items = append(items, fmt.Sprintf("%v as the %v", value.ToLanguage(language), key.ToLanguage(language)))
			return false
		})

		return strings.Join(items, "")
	}

}
