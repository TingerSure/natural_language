package component

import (
	"github.com/TingerSure/natural_language/sandbox/concept"
)

type ClassReflection struct {
	class       concept.Class
	alias       string
	reflections map[string]string
}

func (c *ClassReflection) CheckReflections() bool {
	if len(c.reflections) != len(c.class.AllFields()) {
		return false
	}
	for field, _ := range c.class.AllFields() {
		if c.reflections[field] == "" {
			return false
		}
	}
	return true
}

func (c *ClassReflection) SetReflection(classField string, objectField string) {
	c.reflections[classField] = objectField
}

func (c *ClassReflection) GetReflection(classField string) string {
	return c.reflections[classField]
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
		class:       class,
		reflections: mapping,
		alias:       alias,
	}
}
