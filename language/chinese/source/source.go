package source

import (
	"github.com/TingerSure/natural_language/core/tree"
	"github.com/TingerSure/natural_language/language/chinese/source/priority"
	"github.com/TingerSure/natural_language/language/chinese/source/structs"
	"github.com/TingerSure/natural_language/language/chinese/source/word/auxiliary"
	"github.com/TingerSure/natural_language/language/chinese/source/word/brackets"
	"github.com/TingerSure/natural_language/language/chinese/source/word/number"
	"github.com/TingerSure/natural_language/language/chinese/source/word/operator"
	"github.com/TingerSure/natural_language/language/chinese/source/word/pronoun"
	"github.com/TingerSure/natural_language/language/chinese/source/word/question"
	"github.com/TingerSure/natural_language/language/chinese/source/word/unknown"
	"github.com/TingerSure/natural_language/language/chinese/source/word/verb/set"
)

func AllRules() []tree.Source {
	return []tree.Source{
		pronoun.NewIt(),
		pronoun.NewResult(),
		set.NewSet(),
		question.NewWhat(),
		question.NewHowMany(),
		auxiliary.NewBelong(),
		unknown.NewUnknown(),
		number.NewNumber(),
		operator.NewAddition(),
		operator.NewSubtraction(),
		operator.NewDivision(),
		operator.NewMultiplication(),
		brackets.NewBracketsLeft(),
		brackets.NewBracketsRight(),

		structs.NewAnyFromAnyBelongAny(),
		structs.NewNumberFromNumberOperatorNumber(),
		structs.NewAnyFromBracketAnyBracket(),
		structs.NewAnyFromQuestionSetAny(),
		structs.NewAnyFromAnySetQuestion(),

		priority.NewOperatorLevel(),
		priority.NewBelong(),
	}
}
