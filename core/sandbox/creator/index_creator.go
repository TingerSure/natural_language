package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
)

type IndexCreator struct {
	ConstIndex    *index.ConstIndexCreator
	ResaultIndex  *index.ResaultIndexCreator
	SearchIndex   *index.SearchIndexCreator
	ThisIndex     *index.ThisIndexCreator
	SelfIndex     *index.SelfIndexCreator
	BubbleIndex   *index.BubbleIndexCreator
	LocalIndex    *index.LocalIndexCreator
	ImportIndex   *index.ImportIndexCreator
	PublicIndex   *index.PublicIndexCreator
	PrivateIndex  *index.PrivateIndexCreator
	ProvideIndex  *index.ProvideIndexCreator
	RequireIndex  *index.RequireIndexCreator
	KeyValueIndex *index.KeyValueIndexCreator
	KeyKeyIndex   *index.KeyKeyIndexCreator
	KeyIndex      *index.KeyIndexCreator
}

type IndexCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	ParamCreator     func() concept.Param
	NullCreator      func() concept.Null
	StringCreator    func(string) concept.String
}

func NewIndexCreator(param *IndexCreatorParam) *IndexCreator {
	instance := &IndexCreator{}
	instance.KeyIndex = index.NewKeyIndexCreator(&index.KeyIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
	})
	instance.KeyKeyIndex = index.NewKeyKeyIndexCreator(&index.KeyKeyIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
	})
	instance.KeyValueIndex = index.NewKeyValueIndexCreator(&index.KeyValueIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
	})
	instance.RequireIndex = index.NewRequireIndexCreator(&index.RequireIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ProvideIndex = index.NewProvideIndexCreator(&index.ProvideIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.PrivateIndex = index.NewPrivateIndexCreator(&index.PrivateIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.PublicIndex = index.NewPublicIndexCreator(&index.PublicIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ImportIndex = index.NewImportIndexCreator(&index.ImportIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.LocalIndex = index.NewLocalIndexCreator(&index.LocalIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.BubbleIndex = index.NewBubbleIndexCreator(&index.BubbleIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.SelfIndex = index.NewSelfIndexCreator(&index.SelfIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		StringCreator:    param.StringCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ThisIndex = index.NewThisIndexCreator(&index.ThisIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		StringCreator:    param.StringCreator,
		NullCreator:      param.NullCreator,
	})
	instance.SearchIndex = index.NewSearchIndexCreator(&index.SearchIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ResaultIndex = index.NewResaultIndexCreator(&index.ResaultIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ConstIndex = index.NewConstIndexCreator(&index.ConstIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		ParamCreator:     param.ParamCreator,
		NullCreator:      param.NullCreator,
	})
	return instance
}
