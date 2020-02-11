package question

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

var (
	HowManyContent                  = std.PrintContent
	HowMany        concept.Function = std.Print

	WhatContent                  = std.PrintContent
	What        concept.Function = std.Print
)

func NewQuestion() tree.Page {
	page := tree.NewPageAdaptor()
	page.SetFunction("HowMany", HowMany)
	page.SetConst("HowManyContent", HowManyContent)
	page.SetFunction("What", What)
	page.SetConst("WhatContent", WhatContent)
	return page
}
