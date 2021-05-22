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

func (v *IndexFamily) IsKeyIndex(value concept.Index) (*KeyIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexKeyType {
		index, yes := value.(*KeyIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsKeyKeyIndex(value concept.Index) (*KeyKeyIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexKeyKeyType {
		index, yes := value.(*KeyKeyIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsKeyValueIndex(value concept.Index) (*KeyValueIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexKeyValueType {
		index, yes := value.(*KeyValueIndex)
		return index, yes
	}
	return nil, false
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

func (v *IndexFamily) IsImportIndex(value concept.Index) (*ImportIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexImportType {
		index, yes := value.(*ImportIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsPrivateIndex(value concept.Index) (*PrivateIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexPrivateType {
		index, yes := value.(*PrivateIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsPublicIndex(value concept.Index) (*PublicIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexPublicType {
		index, yes := value.(*PublicIndex)
		return index, yes
	}
	return nil, false
}
func (v *IndexFamily) IsProvideIndex(value concept.Index) (*ProvideIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexProvideType {
		index, yes := value.(*ProvideIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsRequireIndex(value concept.Index) (*RequireIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexRequireType {
		index, yes := value.(*RequireIndex)
		return index, yes
	}
	return nil, false
}
