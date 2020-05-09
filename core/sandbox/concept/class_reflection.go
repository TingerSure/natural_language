package concept

import ()

type ClassReflection struct {
	class   Class
	alias   string
	mapping map[String]String
}

func (c *ClassReflection) SetMapping(mapping map[String]String) {
	c.mapping = mapping
}

func (c *ClassReflection) GetMapping() map[String]String {
	return c.mapping
}

func (c *ClassReflection) GetClass() Class {
	return c.class
}

func (c *ClassReflection) SetAlias(alias string) {
	c.alias = alias
}

func (c *ClassReflection) GetAlias() string {
	return c.alias
}

func NewClassReflection(class Class) *ClassReflection {
	return &ClassReflection{
		class: class,
	}
}

func NewClassReflectionWithMapping(
	class Class,
	mapping map[String]String,
	alias string,
) *ClassReflection {
	return &ClassReflection{
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
}
