package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

const (
	QuestionParam  = "param"
	QuestionResult = "result"
)

type Question struct {
	tree.Page
	output        *std.Std
	HowManyParam  concept.String
	WhatParam     concept.String
	HowManyResult concept.String
	WhatResult    concept.String
}

func (q *Question) HowMany(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	outParam, suspend := q.output.Print(variable.NewParam().Set(q.output.PrintContent, input.Get(q.HowManyParam)), object)
	return variable.NewParam().Set(q.HowManyResult, outParam.Get(q.output.PrintContent)), suspend
}

func (q *Question) What(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	outParam, suspend := q.output.Print(variable.NewParam().Set(q.output.PrintContent, input.Get(q.WhatParam)), object)
	return variable.NewParam().Set(q.WhatResult, outParam.Get(q.output.PrintContent)), suspend
}

func NewQuestion(libs *tree.LibraryManager, output *std.Std) *Question {

	instance := &Question{
		Page:          tree.NewPageAdaptor(),
		output:        output,
		HowManyParam:  variable.NewString(QuestionParam),
		WhatParam:     variable.NewString(QuestionParam),
		HowManyResult: variable.NewString(QuestionResult),
		WhatResult:    variable.NewString(QuestionResult),
	}
	instance.SetFunction(variable.NewString("HowMany"), variable.NewSystemFunction(
		variable.NewString("HowMany"),
		instance.HowMany,
		[]concept.String{
			instance.HowManyParam,
		},
		[]concept.String{
			instance.HowManyResult,
		},
	))

	instance.SetFunction(variable.NewString("What"), variable.NewSystemFunction(
		variable.NewString("What"),
		instance.What,
		[]concept.String{
			instance.WhatParam,
		},
		[]concept.String{
			instance.WhatResult,
		},
	))

	instance.SetConst(variable.NewString("HowManyParam"), instance.HowManyParam)
	instance.SetConst(variable.NewString("WhatParam"), instance.HowManyParam)
	instance.SetConst(variable.NewString("HowManyResult"), instance.HowManyResult)
	instance.SetConst(variable.NewString("WhatResult"), instance.HowManyResult)
	return instance
}
