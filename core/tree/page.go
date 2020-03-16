package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Page interface {
	GetFunction(concept.String) concept.Function
	SetFunction(concept.String, concept.Function) Page
	GetClass(concept.String) concept.Class
	SetClass(concept.String, concept.Class) Page
	GetConst(concept.String) concept.String
	SetConst(concept.String, concept.String) Page
}
