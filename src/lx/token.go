package lx

type TokenType int

const (
	Eof TokenType = iota
	Atom
	Op
	Id
)

type Token struct {
	Type  TokenType
	Value any
}
