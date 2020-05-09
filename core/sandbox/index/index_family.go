package index

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

var (
	IndexFamilyInstance *IndexFamily = newIndexFamily()
)

func newIndexFamily() *IndexFamily {
	return &IndexFamily{}
}

type IndexFamily struct {
}

func (v *IndexFamily) IsConstIndex(value concept.Index) (*ConstIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexConstType {
		index, yes := value.(*ConstIndex)
		return index, yes
	}
	return nil, false
}
