package sandbox

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

func VariableBindLanguage(libs *tree.LibraryManager, language string) {
	variable.StringLanguageSeeds[language] = func(language string, instance *variable.String) string {
		return instance.Value()
	}

	variable.NumberLanguageSeeds[language] = func(language string, instance *variable.Number) string {
		return fmt.Sprintf("%v", instance.Value())
	}

	variable.FunctionLanguageSeeds[language] = func(language string, instance *variable.Function) string {
		return instance.Name().ToLanguage(language)
	}

}
