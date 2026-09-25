package interpreter

import (
	"fmt"
	"rhodium/src/lexer"
	"rhodium/src/parser"
)

func Do(tree *parser.Node) int {
	if tree.Type == parser.Atom {
		i, ok := tree.Value.(int)
		if !ok {
			panic(fmt.Sprintf("Couldn't parse atom %#v", tree.Value))
		}
		return i
	}

	ex, ok := tree.Value.(parser.ExpressionNode)
	if !ok {
		panic(fmt.Sprintf("Couldn't parse node %#v", tree.Value))
	}

	switch ex.Op {
	case lexer.Plus:
		return Do(ex.Left) + Do(ex.Right)
	case lexer.Minus:
		return Do(ex.Left) - Do(ex.Right)
	case lexer.Star:
		return Do(ex.Left) * Do(ex.Right)
	case lexer.Division:
		return Do(ex.Left) / Do(ex.Right)
	case lexer.Modulus:
		return Do(ex.Left) % Do(ex.Right)
	default:
		panic(fmt.Sprintf("Unknown operator: %#v", ex.Op))
	}
}
