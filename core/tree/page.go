package tree

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type Page interface {
	GetFunction(key string) concept.Function
	GetConst(key string) string
	SetFunction(key string, value concept.Function) Page
	SetConst(key string, value string) Page
}
