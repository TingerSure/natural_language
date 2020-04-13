package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

type Question struct {
	tree.Page
	output         *std.Std
	HowManyContent concept.String
	WhatContent    concept.String
}

func (q *Question) HowMany(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return q.output.Print(input, object)
}

func (q *Question) What(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	return q.output.Print(input, object)
}

func NewQuestion(libs *tree.LibraryManager, output *std.Std) *Question {

	instance := &Question{
		Page:           tree.NewPageAdaptor(),
		output:         output,
		HowManyContent: output.PrintContent.Clone(),
		WhatContent:    output.PrintContent.Clone(),
	}
	instance.SetFunction(variable.NewString("HowMany"), variable.NewSystemFunction(
		instance.HowMany,
		[]concept.String{
			instance.HowManyContent,
		},
		[]concept.String{
			instance.HowManyContent,
		},
	))

	instance.SetFunction(variable.NewString("What"), variable.NewSystemFunction(
		instance.What,
		[]concept.String{
			instance.WhatContent,
		},
		[]concept.String{
			instance.WhatContent,
		},
	))
	instance.SetConst(variable.NewString("HowManyContent"), instance.HowManyContent)
	instance.SetConst(variable.NewString("WhatContent"), instance.HowManyContent)
	return instance
}
