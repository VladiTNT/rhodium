package lexer

type TokenType int

const (
	Eof TokenType = iota
	Op
	Num
	Id
)

type Token struct {
	Type  TokenType
	Value any
}
