package sandbox

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/index"
)

func IndexBindLanguage(libs *runtime.LibraryManager, language string) {
	libs.Sandbox.Index.ConstIndex.Seeds[language] = func(language string, instance *index.ConstIndex) string {
		return instance.Value().ToLanguage(language)
	}

	libs.Sandbox.Index.LocalIndex.Seeds[language] = func(language string, instance *index.LocalIndex) string {
		return instance.Key().ToLanguage(language)
	}

	libs.Sandbox.Index.BubbleIndex.Seeds[language] = func(language string, instance *index.BubbleIndex) string {
		return instance.Key().ToLanguage(language)
	}

	libs.Sandbox.Index.SelfIndex.Seeds[language] = func(language string, instance *index.SelfIndex) string {
		return "self"
	}

	libs.Sandbox.Index.ThisIndex.Seeds[language] = func(language string, instance *index.ThisIndex) string {
		return "this"
	}

}
