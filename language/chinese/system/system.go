package system

import (
	"github.com/TingerSure/natural_language/core/tree"
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

func NewSystem(libs *tree.LibraryManager) *tree.PackageAdaptor {

	system := tree.NewPackageAdaptor()
	system.AddSource(pronoun.NewIt(libs))

	system.AddSource(pronoun.NewResult(libs))

	system.AddSource(set.NewSet(libs))

	questionPackage := question.NewQuestion(libs)
	system.AddSource(questionPackage)
	system.AddSource(question.NewWhat(libs, questionPackage))
	system.AddSource(question.NewHowMany(libs, questionPackage))

	system.AddSource(auxiliary.NewBelong(libs))

	system.AddSource(unknown.NewUnknown(libs))

	system.AddSource(number.NewNumber(libs))

	system.AddSource(operator.NewAddition(libs))
	system.AddSource(operator.NewSubtraction(libs))
	system.AddSource(operator.NewDivision(libs))
	system.AddSource(operator.NewMultiplication(libs))

	system.AddSource(brackets.NewBracketsLeft(libs))
	system.AddSource(brackets.NewBracketsRight(libs))

	system.AddSource(structs.NewAnyFromAnyBelongAny(libs))
	system.AddSource(structs.NewNumberFromNumberOperatorNumber(libs))
	system.AddSource(structs.NewAnyFromBracketAnyBracket(libs))
	system.AddSource(structs.NewAnyFromQuestionSetAny(libs, questionPackage))
	system.AddSource(structs.NewAnyFromAnySetQuestion(libs, questionPackage))

	system.AddSource(priority.NewOperatorLevel(libs))
	system.AddSource(priority.NewBelong(libs))
	return system
}