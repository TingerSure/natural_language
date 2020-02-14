package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

type Question struct {
	*tree.PageAdaptor
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

func NewQuestion(output *std.Std) tree.Page {

	instance := &Question{
		PageAdaptor: tree.NewPageAdaptor(),
		output:      output,
	}
	instance.SetFunction("HowMany", variable.NewSystemFunction(
		instance.HowMany,
		[]string{
			HowManyContent,
		},
		[]string{
			HowManyContent,
		},
	))

	instance.SetFunction("What", variable.NewSystemFunction(
		instance.What,
		[]string{
			WhatContent,
		},
		[]string{
			WhatContent,
		},
	))
	instance.SetConst("HowManyContent", HowManyContent)
	instance.SetConst("WhatContent", WhatContent)
	return instance
}
