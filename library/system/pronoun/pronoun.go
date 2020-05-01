package pronoun

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/matcher"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
	"github.com/TingerSure/natural_language/core/tree"
)

type Pronoun struct {
	tree.Page
	It     concept.Index
	Result concept.Index
}

func NewPronoun(libs *tree.LibraryManager) *Pronoun {
	instance := &Pronoun{
		Page: tree.NewPageAdaptor(),
		It: index.NewSearchIndex([]concept.Matcher{
			matcher.NewSystemMatcher(func(concept.Variable) bool {
				return true
			}),
		}),
		Result: index.NewResaultIndex([]concept.Matcher{
			matcher.NewSystemMatcher(func(concept.Variable) bool {
				return true
			}),
		}),
	}

	instance.SetIndex(variable.NewString("It"), instance.It)
	instance.SetIndex(variable.NewString("Result"), instance.Result)

	return instance
}
