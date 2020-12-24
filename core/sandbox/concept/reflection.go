package concept

type Reflection struct {
	class   Class
	alias   string
	mapping map[String]String
}

func (c *Reflection) SetMapping(mapping map[String]String) {
	c.mapping = mapping
}

func (c *Reflection) GetMapping() map[String]String {
	return c.mapping
}

func (c *Reflection) GetClass() Class {
	return c.class
}

func (c *Reflection) SetAlias(alias string) {
	c.alias = alias
}

func (c *Reflection) GetAlias() string {
	return c.alias
}

func NewReflection(class Class) *Reflection {
	return &Reflection{
		class: class,
	}
}

func NewReflectionWithMapping(
	class Class,
	mapping map[String]String,
	alias string,
) *Reflection {
	return &Reflection{
		class:   class,
		mapping: mapping,
		alias:   alias,
	}
}
