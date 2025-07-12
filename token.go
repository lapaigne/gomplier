package main

const (
	UNKNOWN = iota - 2
	EOF
	NEWLINE
	NUMBER
	IDENT
	STRING

	LABEL
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

	EQ
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
