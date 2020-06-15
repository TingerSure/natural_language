package pronoun

import (
	"github.com/TingerSure/natural_language/core/runtime"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/matcher"
	"github.com/TingerSure/natural_language/core/tree"
)

type Pronoun struct {
	tree.Page
	It     concept.Index
	Result concept.Index
}

func NewPronoun(libs *runtime.LibraryManager) *Pronoun {
	instance := &Pronoun{
		Page: tree.NewPageAdaptor(libs.Sandbox),
		It: libs.Sandbox.Index.SearchIndex.New([]concept.Matcher{
			matcher.NewSystemMatcher(func(concept.Variable) bool {
				return true
			}),
		}),
		Result: libs.Sandbox.Index.ResaultIndex.New([]concept.Matcher{
			matcher.NewSystemMatcher(func(concept.Variable) bool {
				return true
			}),
		}),
	}

	instance.SetIndex(libs.Sandbox.Variable.String.New("It"), instance.It)
	instance.SetIndex(libs.Sandbox.Variable.String.New("Result"), instance.Result)

	return instance
}
