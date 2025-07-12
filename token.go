package main

const (
	EOF = iota - 1
	NEWLINE
	NUMBER
	IDENT
	STRING
)

const (
	LABEL = iota + 101
	GOTO
	PRINT
	INPUT
	LET
	IF
	THEN
	ENDIF
	WHILE
	REPEAT
	ENDWHILE
)

const (
	EQ = iota + 201
	PLUS
	MINUS
	ASTERISK
	SLASH
	EQEQ
	NOTEQ
	LT
	LTEQ
	GT
	GTEQ
)

type Token struct {
	text string
	cat  int
}
