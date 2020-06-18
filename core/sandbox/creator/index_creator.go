package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
	"github.com/TingerSure/natural_language/core/sandbox/variable"
)

type IndexCreator struct {
	ConstIndex        *index.ConstIndexCreator
	ObjectFieldIndex  *index.ObjectFieldIndexCreator
	ObjectMethodIndex *index.ObjectMethodIndexCreator
	ResaultIndex      *index.ResaultIndexCreator
	SearchIndex       *index.SearchIndexCreator
	ThisIndex         *index.ThisIndexCreator
	SelfIndex         *index.SelfIndexCreator
	BubbleIndex       *index.BubbleIndexCreator
	LocalIndex        *index.LocalIndexCreator
}

type IndexCreatorParam struct {
	ExceptionCreator         func(string, string) concept.Exception
	NullCreator              func() concept.Null
	StringCreator            func(string) concept.String
	PreObjectFunctionCreator func(concept.Function, concept.Object) *variable.PreObjectFunction
}

func NewIndexCreator(param *IndexCreatorParam) *IndexCreator {
	instance := &IndexCreator{}
	instance.LocalIndex = index.NewLocalIndexCreator(&index.LocalIndexCreatorParam{})
	instance.BubbleIndex = index.NewBubbleIndexCreator(&index.BubbleIndexCreatorParam{})
	instance.SelfIndex = index.NewSelfIndexCreator(&index.SelfIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		StringCreator:    param.StringCreator,
	})
	instance.ThisIndex = index.NewThisIndexCreator(&index.ThisIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		StringCreator:    param.StringCreator,
	})
	instance.SearchIndex = index.NewSearchIndexCreator(&index.SearchIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ResaultIndex = index.NewResaultIndexCreator(&index.ResaultIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ObjectMethodIndex = index.NewObjectMethodIndexCreator(&index.ObjectMethodIndexCreatorParam{
		ExceptionCreator:         param.ExceptionCreator,
		PreObjectFunctionCreator: param.PreObjectFunctionCreator,
		NullCreator:              param.NullCreator,
	})
	instance.ObjectFieldIndex = index.NewObjectFieldIndexCreator(&index.ObjectFieldIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
		NullCreator:      param.NullCreator,
	})
	instance.ConstIndex = index.NewConstIndexCreator(&index.ConstIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	return instance
}
