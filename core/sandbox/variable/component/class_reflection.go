package component

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ClassReflection struct {
	class   concept.Class
	alias   string
	mapping map[concept.String]concept.String
}

func (c *ClassReflection) SetMapping(mapping map[concept.String]concept.String) {
	c.mapping = mapping
}

func (c *ClassReflection) GetMapping() map[concept.String]concept.String {
	return c.mapping
}

func (c *ClassReflection) GetClass() concept.Class {
	return c.class
}

func (c *ClassReflection) SetAlias(alias string) {
	c.alias = alias
}

func (c *ClassReflection) GetAlias() string {
	return c.alias
}

func NewClassReflection(class concept.Class) *ClassReflection {
	return &ClassReflection{
		class: class,
	}
}

func NewClassReflectionWithMapping(
	class concept.Class,
	mapping map[concept.String]concept.String,
	alias string,
) *ClassReflection {
	return &ClassReflection{
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
}
