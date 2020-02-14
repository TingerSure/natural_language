package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Page interface {
	GetClass(key string) concept.Class
	GetFunction(key string) concept.Function
	GetConst(key string) string
	SetClass(key string, value concept.Class) Page
	SetFunction(key string, value concept.Function) Page
	SetConst(key string, value string) Page
}
