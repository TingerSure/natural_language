package lexer

type TokenList struct {
	tokens    []*Token
	index     int
	eof       *Token
	startLine *Line
	trims     map[int]bool
}

func NewTokenList() *TokenList {
	return &TokenList{
		index: 0,
		trims: map[int]bool{},
	}
}

func (t *TokenList) StartLine() *Line {
	return t.startLine
}

func (t *TokenList) SetStartLine(startLine *Line) {
	t.startLine = startLine
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

func (t *TokenList) SetEof(eof *Token) {
	t.eof = eof
}

func (t *TokenList) Eof() *Token {
	return t.eof
}

func (t *TokenList) Reset() {
	t.index = 0
}

func (t *TokenList) Size() int {
	return len(t.tokens)
}

func (t *TokenList) IsEof() bool {
	return t.index > len(t.tokens)
}

func (t *TokenList) Peek() *Token {
	if t.index > len(t.tokens) {
		return nil
	}
	if t.index == len(t.tokens) {
		return t.eof
	}
	return t.tokens[t.index]
}

func (t *TokenList) Next() (token *Token) {
	token = t.Peek()
	t.index++
	return
}
