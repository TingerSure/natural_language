package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/phrase_type"
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
	addPhraseTypes(system, param)
	addWords(system, param)
	addStructs(system, param)
	addPrioritys(system, param)
	return system
}

func addPhraseTypes(system tree.Page, param *adaptor.SourceAdaptorParam) {
	system.AddSource(phrase_type.NewAny(param))
	system.AddSource(phrase_type.NewAuxiliaryBelong(param))
	system.AddSource(phrase_type.NewBool(param))
	system.AddSource(phrase_type.NewBrackets(param))
	system.AddSource(phrase_type.NewEntity(param))
	system.AddSource(phrase_type.NewNoun(param))
	system.AddSource(phrase_type.NewNumber(param))
	system.AddSource(phrase_type.NewOperatorArithmetic(param))
	system.AddSource(phrase_type.NewOperatorLogicalUnary(param))
	system.AddSource(phrase_type.NewOperatorLogical(param))
	system.AddSource(phrase_type.NewOperatorRelational(param))
	system.AddSource(phrase_type.NewPronounInterrogative(param))
	system.AddSource(phrase_type.NewPronounPersonal(param))
	system.AddSource(phrase_type.NewQuestion(param))
	system.AddSource(phrase_type.NewSet(param))
	system.AddSource(phrase_type.NewUnknown(param))
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
	system.AddSource(structs.NewQuestionFromInterrogativeSetEntity(param))
	system.AddSource(structs.NewQuestionFromEntitySetInterrogative(param))
}

func addPrioritys(system tree.Page, param *adaptor.SourceAdaptorParam) {
	system.AddSource(priority.NewOperatorLevel(param))
	system.AddSource(priority.NewBelong(param))
}
