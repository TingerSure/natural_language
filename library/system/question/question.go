package question

import (
	// "github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/library/system/std"
)

// var (
// 	HowManyContent                  = std.PrintContent
// 	HowMany        concept.Function = std.Print
//
// 	WhatContent                  = std.PrintContent
// 	What        concept.Function = std.Print
// )

func NewQuestion(instance *std.Std) tree.Page {
	page := tree.NewPageAdaptor()
	page.SetFunction("HowMany", instance.GetFunction("Print"))
	page.SetConst("HowManyContent", instance.GetConst("PrintContent"))
	page.SetFunction("What", instance.GetFunction("Print"))
	page.SetConst("WhatContent", instance.GetConst("PrintContent"))
	return page
}
