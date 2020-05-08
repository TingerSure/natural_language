package sandbox

import (
	"github.com/TingerSure/natural_language/core/sandbox/expression"
	"github.com/TingerSure/natural_language/core/tree"
)

func ExpressionBindLanguage(libs *tree.LibraryManager, language string) {
	expression.CallLanguageSeeds[language] = func(language string, instance *expression.Call) string {
		return "TODO"
	}

}
