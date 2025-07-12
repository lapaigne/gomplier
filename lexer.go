package main

type Lexer struct {
	src  []rune
	char rune
	pos  int
}

func (l *Lexer) init(src string) {
	l.src = []rune(src + "\n")
	l.char = -1
	l.pos = -1
	l.next()
}

func (l *Lexer) next() {
	l.pos++
	if l.pos >= len(l.src) {
		l.char = rune(0)
	} else {
		l.char = rune(l.src[l.pos])
	}

}

func (l *Lexer) peek() rune {
	if l.pos+1 >= len(l.src) {
		return rune(0)
	}
	return l.src[l.pos+1]
}

func (l *Lexer) abort() error {
	return nil
}

func (l Lexer) skipWSpace() {

}

func (l Lexer) skipComment() {

}

func (l *Lexer) getToken() Token {

	var t Token

	switch l.char {
	case '+':
		t = Token{text: string(l.char), cat: PLUS}
	case '-':
		t = Token{text: string(l.char), cat: MINUS}
	case '/':
		t = Token{text: string(l.char), cat: SLASH}
	case '*':
		t = Token{text: string(l.char), cat: ASTERISK}
	case '\n':
		t = Token{text: string(l.char), cat: NEWLINE}
	case rune(0):
		t = Token{text: string(l.char), cat: EOF}
	default:
	}

	l.next()
	return t
}
