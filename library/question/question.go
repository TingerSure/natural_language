package question

import (
	"github.com/TingerSure/natural_language/library/std"
	"github.com/TingerSure/natural_language/sandbox/concept"
	"github.com/TingerSure/natural_language/sandbox/expression"
	"github.com/TingerSure/natural_language/sandbox/index"
	"github.com/TingerSure/natural_language/sandbox/variable"
)

var (
	WhatFunc    concept.Function = nil
	HowManyFunc concept.Function = nil
)

func init() {

	howManyFunc := variable.NewFunction(nil)
	howManyFunc.AddParamName(std.PrintContent)
	howManyFunc.Body().AddStep(
		expression.NewReturn(
			std.PrintContent,
			expression.NewCall(
				index.NewConstIndex(std.Print),
				expression.NewNewParamWithInit(map[string]concept.Index{
					std.PrintContent: index.NewLocalIndex(std.PrintContent),
				}),
			),
		),
	)

	WhatFunc = std.Print
	HowManyFunc = howManyFunc
}
