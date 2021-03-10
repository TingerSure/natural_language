package lexer

type TokenList struct {
	tokens []*Token
	index  int
	end    *Token
	trims  map[int]bool
}

func NewTokenList() *TokenList {
	return &TokenList{
		index: 0,
		trims: map[int]bool{},
	}
}

func (t *TokenList) AddToken(token *Token) {
	if t.trims[token.Type()] {
		return
	}
	t.tokens = append(t.tokens, token)
}

func (t *TokenList) AddTrims(trims ...int) {
	for _, trim := range trims {
		t.trims[trim] = true
	}
}

func (t *TokenList) SetEnd(end *Token) {
	t.end = end
}

func (t *TokenList) Reset() {
	t.index = 0
}

func (t *TokenList) Size() int {
	return len(t.tokens)
}

func (t *TokenList) IsEnd() bool {
	return t.index >= len(t.tokens)
}

func (t *TokenList) Next() (token *Token) {
	if t.index >= len(t.tokens) {
		token = t.end
		return
	}
	token = t.tokens[t.index]
	t.index++
	return
}
