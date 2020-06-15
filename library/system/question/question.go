package question

import (
	"github.com/TingerSure/natural_language/core/runtime"
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
	outParam, suspend := q.output.Print(libs.Sandbox.Variable.Param.New().Set(q.output.PrintContent, input.Get(q.HowManyParam)), object)
	return libs.Sandbox.Variable.Param.New().Set(q.HowManyResult, outParam.Get(q.output.PrintContent)), suspend
}

func (q *Question) What(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
	outParam, suspend := q.output.Print(libs.Sandbox.Variable.Param.New().Set(q.output.PrintContent, input.Get(q.WhatParam)), object)
	return libs.Sandbox.Variable.Param.New().Set(q.WhatResult, outParam.Get(q.output.PrintContent)), suspend
}

func NewQuestion(libs *runtime.LibraryManager, output *std.Std) *Question {

	instance := &Question{
		Page:          tree.NewPageAdaptor(libs.Sandbox),
		output:        output,
		HowManyParam:  libs.Sandbox.Variable.String.New(QuestionParam),
		WhatParam:     libs.Sandbox.Variable.String.New(QuestionParam),
		HowManyResult: libs.Sandbox.Variable.String.New(QuestionResult),
		WhatResult:    libs.Sandbox.Variable.String.New(QuestionResult),
	}
	instance.SetFunction(libs.Sandbox.Variable.String.New("HowMany"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("HowMany"),
		instance.HowMany,
		[]concept.String{
			instance.HowManyParam,
		},
		[]concept.String{
			instance.HowManyResult,
		},
	))

	instance.SetFunction(libs.Sandbox.Variable.String.New("What"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("What"),
		instance.What,
		[]concept.String{
			instance.WhatParam,
		},
		[]concept.String{
			instance.WhatResult,
		},
	))

	instance.SetConst(libs.Sandbox.Variable.String.New("HowManyParam"), instance.HowManyParam)
	instance.SetConst(libs.Sandbox.Variable.String.New("WhatParam"), instance.HowManyParam)
	instance.SetConst(libs.Sandbox.Variable.String.New("HowManyResult"), instance.HowManyResult)
	instance.SetConst(libs.Sandbox.Variable.String.New("WhatResult"), instance.HowManyResult)
	return instance
}
