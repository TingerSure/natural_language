package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Page interface {
	concept.Index
	GetFunction(concept.String) concept.Function
	SetFunction(concept.String, concept.Function) Page
	GetClass(concept.String) concept.Class
	SetClass(concept.String, concept.Class) Page
	GetConst(concept.String) concept.String
	SetConst(concept.String, concept.String) Page
	GetIndex(concept.String) concept.Index
	SetIndex(concept.String, concept.Index) Page
	GetException(concept.String) concept.Exception
	SetException(concept.String, concept.Exception) Page
	GetSources() []Source
	AddSource(Source)
}
