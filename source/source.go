package source

import (
	"github.com/TingerSure/natural_language/source/priority"
	"github.com/TingerSure/natural_language/source/structs"
	"github.com/TingerSure/natural_language/source/word/auxiliary"
	"github.com/TingerSure/natural_language/source/word/brackets"
	"github.com/TingerSure/natural_language/source/word/number"
	"github.com/TingerSure/natural_language/source/word/operator"
	"github.com/TingerSure/natural_language/source/word/pronoun"
	"github.com/TingerSure/natural_language/source/word/question"
	"github.com/TingerSure/natural_language/source/word/unknown"
	"github.com/TingerSure/natural_language/source/word/verb/set"
	"github.com/TingerSure/natural_language/tree"
)

func AllRules() []tree.Source {
	return []tree.Source{
		pronoun.NewTarget(),
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

		structs.NewTargetFromTargetTarget(),
		structs.NewTargetFromTargetBelongTarget(),
		structs.NewEventFromTargetActionTarget(),
		structs.NewNumberFromNumberOperatorNumber(),
		structs.NewAnyFromBracketAnyBracket(),
		structs.NewAnyFromQuestionSetAny(),
		structs.NewAnyFromAnySetQuestion(),

		priority.NewOperatorLevel(),
	}
}
