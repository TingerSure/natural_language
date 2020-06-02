package creator

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/sandbox/index"
)

type InterruptCreator struct {
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

type InterruptCreatorParam struct {
	ExceptionCreator func(string, string) concept.Exception
	NullCreator      func() concept.Null
	StringCreator    func(string) concept.String
}

func NewInterruptCreator(param *InterruptCreatorParam) *InterruptCreator {
	instance := &InterruptCreator{}
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
	instance.ObjectMethodIndex = index.NewObjectMethodIndexCreator(&index.ObjectFieldIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.ObjectFieldIndex = index.NewObjectFieldIndexCreator(&index.ObjectFieldIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	instance.ConstIndex = index.NewConstIndexCreator(&index.ConstIndexCreatorParam{
		ExceptionCreator: param.ExceptionCreator,
	})
	return instance
}
