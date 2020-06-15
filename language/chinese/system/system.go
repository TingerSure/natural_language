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

	system := tree.NewPageAdaptor(libs.Sandbox)
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

	system.AddSource(structs.NewAnyFromAnyBelongAny(param))
	system.AddSource(structs.NewNumberFromNumberOperatorNumber(param))
	system.AddSource(structs.NewAnyFromBracketAnyBracket(param))
	system.AddSource(structs.NewAnyFromQuestionSetAny(param))
	system.AddSource(structs.NewAnyFromAnySetQuestion(param))

	system.AddSource(priority.NewOperatorLevel(param))
	system.AddSource(priority.NewBelong(param))
	return system
}
