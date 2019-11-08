package source

import (
	"github.com/TingerSure/natural_language/source/rule"
	"github.com/TingerSure/natural_language/source/system/auxiliary"
	"github.com/TingerSure/natural_language/source/system/number"
	"github.com/TingerSure/natural_language/source/system/operator"
	"github.com/TingerSure/natural_language/source/system/pronoun"
	"github.com/TingerSure/natural_language/source/system/unknown"
	"github.com/TingerSure/natural_language/source/system/verb"
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

		rule.NewTargetFromTargetTarget(),
		rule.NewTargetFromUnknown(),
		rule.NewTargetFromTargetBelongTarget(),
		rule.NewEventFromTargetActionTarget(),
		rule.NewNumberFromNumberOperatorNumber(),
	}
}
