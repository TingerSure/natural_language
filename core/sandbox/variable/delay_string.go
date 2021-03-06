package variable

import (
	"github.com/TingerSure/natural_language/core/adaptor/nl_interface"
	"github.com/TingerSure/natural_language/core/sandbox/concept"
)

const (
	VariableDelayStringType = "delay_string"
)

type DelayStringSeed interface {
	ToLanguage(string, concept.Pool, *DelayString) (string, concept.Exception)
	Type() string
	NewString(string) concept.String
}

type DelayString struct {
	original string
	value    concept.String
	seed     DelayStringSeed
}

func (o *DelayString) init() {
	if nl_interface.IsNil(o.value) {
		o.value = o.seed.NewString(o.original)
	}
}

func (o *DelayString) IsFunction() bool {
	return false
}

func (o *DelayString) IsNull() bool {
	return false
}

func (o *DelayString) SetField(specimen concept.String, value concept.Variable) concept.Exception {
	o.init()
	return o.value.SetField(specimen, value)
}

func (o *DelayString) GetField(specimen concept.String) (concept.Variable, concept.Exception) {
	o.init()
	return o.value.GetField(specimen)
}

func (o *DelayString) HasField(specimen concept.String) bool {
	o.init()
	return o.value.HasField(specimen)
}

func (o *DelayString) KeyField(specimen concept.String) concept.String {
	o.init()
	return o.value.KeyField(specimen)
}

func (o *DelayString) SizeField() int {
	o.init()
	return o.value.SizeField()
}

func (o *DelayString) Iterate(on func(concept.String, concept.Variable) bool) bool {
	o.init()
	return o.value.Iterate(on)
}

func (o *DelayString) ToString(prefix string) string {
	o.init()
	return o.value.ToString(prefix)
}

func (o *DelayString) Call(specimen concept.String, param concept.Param) (concept.Param, concept.Exception) {
	o.init()
	return o.value.Call(specimen, param)
}

func (f *DelayString) ToLanguage(language string, space concept.Pool) (string, concept.Exception) {
	return f.seed.ToLanguage(language, space, f)
}

func (n *DelayString) GetLanguage(language string) string {
	n.init()
	return n.value.GetLanguage(language)
}

func (n *DelayString) SetLanguage(language string, value string) {
	n.init()
	n.value.SetLanguage(language, value)
}

func (n *DelayString) HasLanguage(language string) bool {
	n.init()
	return n.value.HasLanguage(language)
}

func (n *DelayString) IsLanguage(language string, value string) bool {
	n.init()
	return n.value.IsLanguage(language, value)
}

func (n *DelayString) Equal(other concept.String) bool {
	n.init()
	return n.value.Equal(other)
}

func (n *DelayString) IterateLanguages(on func(string, string) bool) bool {
	n.init()
	return n.value.IterateLanguages(on)
}

func (n *DelayString) Value() string {
	return n.original
}

func (n *DelayString) MapKey() string {
	return n.original
}

func (s *DelayString) Type() string {
	return s.seed.Type()
}

func (n *DelayString) Clone() concept.String {
	n.init()
	return n.value.Clone()
}

func (n *DelayString) CloneTo(instance concept.String) {
	n.init()
	n.value.CloneTo(instance)
}

type DelayStringCreatorParam struct {
	StringCreator func(string) concept.String
}

type DelayStringCreator struct {
	Seeds map[string]func(concept.Pool, *DelayString) (string, concept.Exception)
	param *DelayStringCreatorParam
}

func (s *DelayStringCreator) New(original string) *DelayString {
	return &DelayString{
		original: original,
		seed:     s,
	}
}

func (s *DelayStringCreator) NewString(value string) concept.String {
	return s.param.StringCreator(value)
}

func (s *DelayStringCreator) ToLanguage(language string, space concept.Pool, instance *DelayString) (string, concept.Exception) {
	seed := s.Seeds[language]
	if seed == nil {
		return instance.ToString(""), nil
	}
	return seed(space, instance)
}

func (s *DelayStringCreator) Type() string {
	return VariableDelayStringType
}

func NewDelayStringCreator(param *DelayStringCreatorParam) *DelayStringCreator {
	return &DelayStringCreator{
		Seeds: map[string]func(concept.Pool, *DelayString) (string, concept.Exception){},
		param: param,
	}
}
