package question

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
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
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			outParam, suspend := instance.output.Print(libs.Sandbox.Variable.Param.New().Set(instance.output.PrintContent, input.Get(instance.HowManyParam)), object)
			return libs.Sandbox.Variable.Param.New().Set(instance.HowManyResult, outParam.Get(instance.output.PrintContent)), suspend
		},
		func(input concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(instance.HowManyResult, input.Get(instance.output.PrintContent))
		},
		[]concept.String{
			instance.HowManyParam,
		},
		[]concept.String{
			instance.HowManyResult,
		},
	))

	instance.SetFunction(libs.Sandbox.Variable.String.New("What"), libs.Sandbox.Variable.SystemFunction.New(
		libs.Sandbox.Variable.String.New("What"),
		func(input concept.Param, object concept.Object) (concept.Param, concept.Exception) {
			outParam, suspend := instance.output.Print(libs.Sandbox.Variable.Param.New().Set(instance.output.PrintContent, input.Get(instance.WhatParam)), object)
			return libs.Sandbox.Variable.Param.New().Set(instance.WhatResult, outParam.Get(instance.output.PrintContent)), suspend
		},
		func(input concept.Param, _ concept.Object) concept.Param {
			return libs.Sandbox.Variable.Param.New().Set(instance.WhatResult, input.Get(instance.output.PrintContent))
		},
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
