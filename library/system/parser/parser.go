package parser

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
	"github.com/TingerSure/natural_language/core/tree"
)

type Parser struct {
	concept.Page
	libs *tree.LibraryManager
}

func NewParser(libs *tree.LibraryManager) *Parser {
	instance := &Parser{
		libs: libs,
		Page: libs.Sandbox.Variable.Page.New(),
	}

	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("addTypes"),
		libs.Sandbox.Index.PublicIndex.New(
			"addTypes",
			libs.Sandbox.Index.ConstIndex.New(newAddTypes(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("addVocabularyRule"),
		libs.Sandbox.Index.PublicIndex.New(
			"addVocabularyRule",
			libs.Sandbox.Index.ConstIndex.New(newAddVocabularyRule(libs)),
		),
	)

	instance.SetPublic(
		libs.Sandbox.Variable.DelayString.New("addStructRule"),
		libs.Sandbox.Index.PublicIndex.New(
			"addStructRule",
			libs.Sandbox.Index.ConstIndex.New(newAddStructRule(libs)),
		),
	)

	return instance
}
