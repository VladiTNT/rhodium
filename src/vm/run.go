package vm

import (
	"fmt"
	"rhodium/src/ps"
)

func Run(tree *ps.Node) *ps.Node {
	if tree.Type == ps.Atom {
		return tree
	}

	ex, ok := tree.Value.(ps.NodeExpression)
	if !ok {
		panic(fmt.Sprintf("Failed to cast into ps.NodeExpression: %v\n", tree.Value))
	}

	// Parse the rest of the tree
	valueLeft := Run(ex.Left).Value
	valueRight := Run(ex.Right).Value

	atomLeft, ok := valueLeft.(ps.NodeAtom)
	if !ok {
		panic(fmt.Sprintf("Failed to cast into ps.NodeAtom: %v\n", valueLeft))
	}

	atomRight, ok := valueRight.(ps.NodeAtom)
	if !ok {
		panic(fmt.Sprintf("Failed to cast into ps.NodeAtom: %v\n", valueRight))
	}

	return &ps.Node{
		Type:  ps.Atom,
		Value: SolveAtoms(atomLeft, atomRight, ex.Op),
	}
}

func SolveAtoms(atomLeft, atomRight ps.NodeAtom, op string) ps.NodeAtom
