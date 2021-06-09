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

	return instance
}