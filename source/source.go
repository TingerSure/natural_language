package source

import (
	"github.com/TingerSure/natural_language/source/priority"
	"github.com/TingerSure/natural_language/source/structs"
	"github.com/TingerSure/natural_language/source/word/auxiliary"
	"github.com/TingerSure/natural_language/source/word/number"
	"github.com/TingerSure/natural_language/source/word/operator"
	"github.com/TingerSure/natural_language/source/word/pronoun"
	"github.com/TingerSure/natural_language/source/word/unknown"
	"github.com/TingerSure/natural_language/source/word/verb"
	"github.com/TingerSure/natural_language/tree"
)

func AllRules() []tree.Source {
	return []tree.Source{
		pronoun.NewTarget(),
		verb.NewSet(),
		auxiliary.NewBelong(),
		unknown.NewUnknown(),
		number.NewNumber(),
		operator.NewAddition(),
		operator.NewSubtraction(),
		operator.NewDivision(),
		operator.NewMultiplication(),

		structs.NewTargetFromTargetTarget(),
		structs.NewTargetFromUnknown(),
		structs.NewTargetFromTargetBelongTarget(),
		structs.NewEventFromTargetActionTarget(),
		structs.NewNumberFromNumberOperatorNumber(),

		priority.NewOperatorLevel(),
	}
}
