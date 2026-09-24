package parser

import "rhodium/src/lexer"

type NodeType int

const (
	Atom NodeType = iota
	Expression
)

type Node struct {
	Type  NodeType
	Value any
}

type ExpressionNode struct {
	Op    lexer.Operator
	Left  *Node
	Right *Node
}
