package sandbox

import (
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/tree"
)

func IndexBindLanguage(libs *tree.LibraryManager, language string) {
	index.ConstIndexLanguageSeeds[language] = func(language string, instance *index.ConstIndex) string {
		return instance.Value().ToLanguage(language)
	}

}
