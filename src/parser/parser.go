package parser

import (
	"fmt"
	"rhodium/src/lexer"
)

func PrattParse(tokens []lexer.Token, index int, minBp float64) (*Node, int) {
	// Next token
	leftToken := tokens[index]
	index++

	// If the token isn't a number we panic
	if leftToken.Type != lexer.Num {
		panic(fmt.Sprintf("Bad token: %#v", leftToken))
	}

	leftNode := &Node{
		Type:  Atom,
		Value: leftToken.Value,
	}

loop:
	for {
		// Peek at the next token
		opToken := tokens[index]

		// Handle the expression based on it's type
		switch opToken.Type {
		// Break out of the loop if we find Eof
		case lexer.Eof:
			break loop
		case lexer.Op:

		default:
			panic(fmt.Sprintf("Bad token: %#v", opToken))
		}

		op, ok := opToken.Value.(lexer.Operator)
		if !ok {
			panic(fmt.Sprintf("Type cast of op token %#v failed", opToken))
		}

		lBp, rBp := op.BindingPower()
		if lBp < minBp {
			break loop
		}

		// We consume the op token
		index++

		// Parse right side
		var rightNode *Node
		rightNode, index = PrattParse(tokens, index, rBp)

		leftNode = &Node{
			Type: Expression,
			Value: ExpressionNode{
				Op:    op,
				Left:  leftNode,
				Right: rightNode,
			},
		}
	}

	return leftNode, index
}
