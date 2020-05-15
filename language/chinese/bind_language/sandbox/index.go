package sandbox

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/index"
)

func IndexBindLanguage(libs *runtime.LibraryManager, language string) {
	index.ConstIndexLanguageSeeds[language] = func(language string, instance *index.ConstIndex) string {
		return instance.Value().ToLanguage(language)
	}

	index.LocalIndexLanguageSeeds[language] = func(language string, instance *index.LocalIndex) string {
		return instance.Key().ToLanguage(language)
	}

	index.BubbleIndexLanguageSeeds[language] = func(language string, instance *index.BubbleIndex) string {
		return instance.Key().ToLanguage(language)
	}

	index.SelfIndexLanguageSeeds[language] = func(language string, instance *index.SelfIndex) string {
		return "自己"
	}

	index.ThisIndexLanguageSeeds[language] = func(language string, instance *index.ThisIndex) string {
		return "这"
	}

}
