package system

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/system/adaptor"
	"github.com/TingerSure/natural_language/language/chinese/system/duty"
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

func BindRule(libs *tree.LibraryManager, chineseName string) {
	param := &adaptor.SourceAdaptorParam{
		Libs:     libs,
		Language: chineseName,
	}
	addPhraseTypes(libs, param)
	addWords(libs, param)
	addStructs(libs, param)
	addDuties(libs, param)
	addPrioritys(libs, param)
}

func addDuties(libs *tree.LibraryManager, param *adaptor.SourceAdaptorParam) {
	libs.AddSource(duty.NewNumber(param))
}

func addPhraseTypes(libs *tree.LibraryManager, param *adaptor.SourceAdaptorParam) {
	libs.AddSource(phrase_type.NewAny(param))
	libs.AddSource(phrase_type.NewAuxiliaryBelong(param))
	libs.AddSource(phrase_type.NewBool(param))
	libs.AddSource(phrase_type.NewBrackets(param))
	libs.AddSource(phrase_type.NewEntity(param))
	libs.AddSource(phrase_type.NewNoun(param))
	libs.AddSource(phrase_type.NewNumber(param))
	libs.AddSource(phrase_type.NewOperatorArithmetic(param))
	libs.AddSource(phrase_type.NewOperatorLogicalUnary(param))
	libs.AddSource(phrase_type.NewOperatorLogical(param))
	libs.AddSource(phrase_type.NewOperatorRelational(param))
	libs.AddSource(phrase_type.NewPronounInterrogative(param))
	libs.AddSource(phrase_type.NewPronounPersonal(param))
	libs.AddSource(phrase_type.NewQuestion(param))
	libs.AddSource(phrase_type.NewSet(param))
	libs.AddSource(phrase_type.NewUnknown(param))
}

func addWords(libs *tree.LibraryManager, param *adaptor.SourceAdaptorParam) {

	libs.AddSource(pronoun.NewIt(param))

	libs.AddSource(pronoun.NewResult(param))

	libs.AddSource(set.NewSet(param))

	libs.AddSource(question.NewWhat(param))
	libs.AddSource(question.NewHowMany(param))

	libs.AddSource(auxiliary.NewBelong(param))

	libs.AddSource(unknown.NewUnknown(param))

	libs.AddSource(number.NewNumber(param))

	libs.AddSource(operator.NewAddition(param))
	libs.AddSource(operator.NewSubtraction(param))
	libs.AddSource(operator.NewDivision(param))
	libs.AddSource(operator.NewMultiplication(param))

	libs.AddSource(brackets.NewBracketsLeft(param))
	libs.AddSource(brackets.NewBracketsRight(param))
}

func addStructs(libs *tree.LibraryManager, param *adaptor.SourceAdaptorParam) {
	libs.AddSource(structs.NewAnticipateFromEntityBelongNoun(param))
	libs.AddSource(structs.NewAnticipateFromEntityBelongUnknown(param))
	libs.AddSource(structs.NewNumberFromNumberArithmeticNumber(param))
	libs.AddSource(structs.NewBoolFromLogicalBool(param))
	libs.AddSource(structs.NewBoolFromBoolLogicalBool(param))
	libs.AddSource(structs.NewBoolFromNumberRelationalNumber(param))
	libs.AddSource(structs.NewDynamicFromBracketAnyBracket(param))
	libs.AddSource(structs.NewQuestionFromInterrogativeSetEntity(param))
	libs.AddSource(structs.NewQuestionFromEntitySetInterrogative(param))
}

func addPrioritys(libs *tree.LibraryManager, param *adaptor.SourceAdaptorParam) {
	libs.AddSource(priority.NewOperatorLevel(param))
	libs.AddSource(priority.NewBelong(param))
}
