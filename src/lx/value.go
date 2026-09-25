package lx

type ValueType int

const (
	Integer ValueType = iota
	Float
	String
)

type Value struct {
	Type  ValueType
	Value string
}
