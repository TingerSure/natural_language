package bind

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Bind struct {
	concept.Page
	libs *tree.LibraryManager
}

func NewBind(libs *tree.LibraryManager) *Bind {
	instance := &Bind{
		libs: libs,
		Page: libs.Sandbox.Variable.Page.New(),
	}

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("valueLanguage"),
		libs.Sandbox.Index.PublicIndex.New(
			"valueLanguage",
			libs.Sandbox.Index.ConstIndex.New(newValueLanguage(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("keyBind"),
		libs.Sandbox.Index.PublicIndex.New(
			"keyBind",
			libs.Sandbox.Index.ConstIndex.New(newKeyBind(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("codeBlockBind"),
		libs.Sandbox.Index.PublicIndex.New(
			"codeBlockBind",
			libs.Sandbox.Index.ConstIndex.New(newCodeBlockBind(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("bubbleIndexBind"),
		libs.Sandbox.Index.PublicIndex.New(
			"bubbleIndexBind",
			libs.Sandbox.Index.ConstIndex.New(newBubbleIndexBind(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("constIndexBind"),
		libs.Sandbox.Index.PublicIndex.New(
			"constIndexBind",
			libs.Sandbox.Index.ConstIndex.New(newConstIndexBind(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.String.New("stringBind"),
		libs.Sandbox.Index.PublicIndex.New(
			"stringBind",
			libs.Sandbox.Index.ConstIndex.New(newStringBind(libs)),
		),
	)

	return instance
}
