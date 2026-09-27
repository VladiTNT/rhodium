package ps

type NodeType int

const (
	Atom NodeType = iota
	Expression
)

type Node struct {
	Type  NodeType
	Value any
}
