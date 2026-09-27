package ps

import (
	"fmt"
	"rhodium/src/atom"
	"strconv"
)

type NodeAtom struct {
	Type  atom.Type
	Value string
}

func (na *NodeAtom) GetInteger() int {
	i, err := strconv.Atoi(na.Value)
	if err != nil {
		panic(fmt.Sprintf("Parsing '%s' as integer failed.\n", na.Value))
	}

	return i
}

func (na *NodeAtom) GetFloat() float64 {
	f, err := strconv.ParseFloat(na.Value, 64)
	if err != nil {
		panic(fmt.Sprintf("Parsing '%s' as float failed.\n", na.Value))
	}

	return f
}

func (na *NodeAtom) GetString() string {
	return na.Value
}

func (na *NodeAtom) GetValue() any {
	switch na.Type {
	case atom.Integer:
		return na.GetInteger()
	case atom.Float:
		return na.GetFloat()
	case atom.String:
		return na.GetString()
	}
	panic("How does this even happen?")
}
