package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

func QuestionBindLanguage(libs *runtime.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "question")

	HowManyParam := instance.GetConst(variable.NewString("HowManyParam"))
	HowManyResult := instance.GetConst(variable.NewString("HowManyResult"))
	HowMany := instance.GetFunction(variable.NewString("HowMany"))

	HowManyParam.SetLanguage(language, "content")
	HowManyResult.SetLanguage(language, "result")
	HowMany.Name().SetLanguage(language, "how many")
	HowMany.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(HowManyParam).(concept.ToString)
		return fmt.Sprintf("how many are %v", content.ToLanguage(language))

	})

	WhatParam := instance.GetConst(variable.NewString("WhatParam"))
	WhatResult := instance.GetConst(variable.NewString("WhatResult"))
	What := instance.GetFunction(variable.NewString("What"))

	WhatParam.SetLanguage(language, "content")
	WhatResult.SetLanguage(language, "result")
	What.Name().SetLanguage(language, "what")
	What.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(WhatParam).(concept.ToString)
		return fmt.Sprintf("what is %v", content.ToLanguage(language))

	})
}
