package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

type Question struct {
	tree.Page
	output *std.Std
}

var (
	HowManyContent = std.PrintContent
	WhatContent    = std.PrintContent
)

func (q *Question) HowMany(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return q.output.Print(input, object)
}

func (q *Question) What(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return q.output.Print(input, object)
}

func NewQuestion(output *std.Std) *Question {

	instance := &Question{
		Page:   tree.NewPageAdaptor(),
		output: output,
	}
	instance.SetFunction(variable.NewString("HowMany"), variable.NewSystemFunction(
		instance.HowMany,
		[]concept.String{
			variable.NewString(HowManyContent),
		},
		[]concept.String{
			variable.NewString(HowManyContent),
		},
	))

	instance.SetFunction(variable.NewString("What"), variable.NewSystemFunction(
		instance.What,
		[]concept.String{
			variable.NewString(WhatContent),
		},
		[]concept.String{
			variable.NewString(WhatContent),
		},
	))
	instance.SetConst(variable.NewString("HowManyContent"), variable.NewString(HowManyContent))
	instance.SetConst(variable.NewString("WhatContent"), variable.NewString(WhatContent))
	return instance
}
