package component

import (
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

type ClassReflection struct {
	class   concept.Class
	alias   string
	mapping map[string]string
}

func (c *ClassReflection) CheckMapping() bool {
	if len(c.mapping) != len(c.class.AllFields()) {
		return false
	}
	for field, _ := range c.class.AllFields() {
		if c.mapping[field] == "" {
			return false
		}
	}
	return true
}

func (c *ClassReflection) SetMapping(mapping map[string]string) {
	c.mapping = mapping
}

func (c *ClassReflection) GetMapping() map[string]string {
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

func NewClassReflectionWithMapping(class concept.Class, mapping map[string]string, alias string) *ClassReflection {
	return &ClassReflection{
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
}
