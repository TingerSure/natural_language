package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/priority"
	"github.com/TingerSure/natural_language/language/chinese/system/structs"
	"github.com/TingerSure/natural_language/language/chinese/system/word/auxiliary"
	"github.com/TingerSure/natural_language/language/chinese/system/word/brackets"
	"github.com/TingerSure/natural_language/language/chinese/system/word/number"
	"github.com/TingerSure/natural_language/language/chinese/system/word/operator"
	"github.com/TingerSure/natural_language/language/chinese/system/word/pronoun"
	"github.com/TingerSure/natural_language/language/chinese/system/word/question"
	"github.com/TingerSure/natural_language/language/chinese/system/word/unknown"
	"github.com/TingerSure/natural_language/language/chinese/system/word/verb/set"
)

func NewSystem(param *adaptor.SourceAdaptorParam) tree.Page {

	system := tree.NewPageAdaptor(param.Libs.Sandbox)
	addWords(system, param)
	addStructs(system, param)
	addPrioritys(system, param)
	return system
}

func addWords(system tree.Page, param *adaptor.SourceAdaptorParam) {

	system.AddSource(pronoun.NewIt(param))

	system.AddSource(pronoun.NewResult(param))

	system.AddSource(set.NewSet(param))

	system.AddSource(question.NewWhat(param))
	system.AddSource(question.NewHowMany(param))

	system.AddSource(auxiliary.NewBelong(param))

	system.AddSource(unknown.NewUnknown(param))

	system.AddSource(number.NewNumber(param))

	system.AddSource(operator.NewAddition(param))
	system.AddSource(operator.NewSubtraction(param))
	system.AddSource(operator.NewDivision(param))
	system.AddSource(operator.NewMultiplication(param))

	system.AddSource(brackets.NewBracketsLeft(param))
	system.AddSource(brackets.NewBracketsRight(param))
}

func addStructs(system tree.Page, param *adaptor.SourceAdaptorParam) {
	system.AddSource(structs.NewEntityFromEntityBelongNoun(param))
	system.AddSource(structs.NewNumberFromNumberArithmeticNumber(param))
	system.AddSource(structs.NewBoolFromLogicalBool(param))
	system.AddSource(structs.NewBoolFromBoolLogicalBool(param))
	system.AddSource(structs.NewBoolFromNumberRelationalNumber(param))
	system.AddSource(structs.NewAnyFromBracketAnyBracket(param))
	system.AddSource(structs.NewEntityFromNoun(param))
	system.AddSource(structs.NewEntityFromNumber(param))
	system.AddSource(structs.NewQuestionFromInterrogativeSetEntity(param))
	system.AddSource(structs.NewQuestionFromEntitySetInterrogative(param))
}

func addPrioritys(system tree.Page, param *adaptor.SourceAdaptorParam) {
	system.AddSource(priority.NewOperatorLevel(param))
	system.AddSource(priority.NewBelong(param))
}
