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

	HowManyParam.SetLanguage(language, "内容")
	HowManyResult.SetLanguage(language, "结果")
	HowMany.Name().SetLanguage(language, "展示数量")
	HowMany.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(HowManyParam).(concept.ToString)
		return fmt.Sprintf("%v是多少", content.ToLanguage(language))

	})

	WhatParam := instance.GetConst(variable.NewString("WhatParam"))
	WhatResult := instance.GetConst(variable.NewString("WhatResult"))
	What := instance.GetFunction(variable.NewString("What"))

	WhatParam.SetLanguage(language, "内容")
	WhatResult.SetLanguage(language, "结果")
	What.Name().SetLanguage(language, "展示内容")
	What.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *concept.Mapping) string {
		content := param.Get(WhatParam).(concept.ToString)
		return fmt.Sprintf("%v是什么", content.ToLanguage(language))

	})
}
