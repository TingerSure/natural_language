package lexer

import (
	"regexp"
)

type Rule struct {
	matcher *regexp.Regexp
	create  func(value []byte) *Token
}

func NewRule(template string, create func(value []byte) *Token) *Rule {
	matcher := regexp.MustCompile("^" + template)
	return &Rule{
		matcher: matcher,
		create:  create,
	}
}

func (r *Rule) GetMatcher() *regexp.Regexp {
	return r.matcher
}

func (r *Rule) CreateToken(value []byte) *Token {
	return r.create(value)
}
