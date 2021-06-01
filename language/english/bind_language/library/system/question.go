package system

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

func QuestionBindLanguage(libs *tree.LibraryManager, language string) {
	instance := libs.GetLibraryPage("system", "question")

	HowManyParam := instance.GetConst(libs.Sandbox.Variable.String.New("HowManyParam"))
	HowManyResult := instance.GetConst(libs.Sandbox.Variable.String.New("HowManyResult"))
	HowMany := instance.GetFunction(libs.Sandbox.Variable.String.New("HowMany"))

	HowManyParam.SetLanguage(language, "content")
	HowManyResult.SetLanguage(language, "result")
	HowMany.Name().SetLanguage(language, "how many")
	HowMany.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *nl_interface.Mapping) string {
		content := param.Get(HowManyParam).(concept.ToString)
		return fmt.Sprintf("how many are %v", content.ToLanguage(language))

	})

	WhatParam := instance.GetConst(libs.Sandbox.Variable.String.New("WhatParam"))
	WhatResult := instance.GetConst(libs.Sandbox.Variable.String.New("WhatResult"))
	What := instance.GetFunction(libs.Sandbox.Variable.String.New("What"))

	WhatParam.SetLanguage(language, "content")
	WhatResult.SetLanguage(language, "result")
	What.Name().SetLanguage(language, "what")
	What.SetLanguageOnCallSeed(language, func(funcs concept.Function, param *nl_interface.Mapping) string {
		content := param.Get(WhatParam).(concept.ToString)
		return fmt.Sprintf("what is %v", content.ToLanguage(language))

	})
}
