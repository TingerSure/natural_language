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

func (v *IndexFamily) IsVarIndex(value concept.Index) (*VarIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexVarType {
		index, yes := value.(*VarIndex)
		return index, yes
	}
	return nil, false
}

func (v *IndexFamily) IsExportIndex(value concept.Index) (*ExportIndex, bool) {
	if value == nil {
		return nil, false
	}
	if value.Type() == IndexExportType {
		index, yes := value.(*ExportIndex)
		return index, yes
	}
	return nil, false
}
