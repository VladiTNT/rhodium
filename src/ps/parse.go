package ps

import (
	"fmt"
	"rhodium/src/lx"
)

func ParseExpression(tokens []lx.Token, index *int, minBp float64) *Node {
	// Get left side token
	leftToken := tokens[*index]
	*index++

	// Token must be an atom
	if leftToken.Type != lx.Atom {
		panic(fmt.Sprintf("Bad token, should have been atom: %v\n", leftToken))
	}

	// Get the value of this token to put in node
	leftValue, ok := leftToken.Value.(lx.Value)
	if !ok {
		panic(fmt.Sprintf("Type cast to lx.Value failed: %v\n", leftToken.Value))
	}

	leftNode := &Node{
		Type: Atom,
		Value: NodeAtom{
			Type:  leftValue.Type,
			Value: leftValue.Value,
		},
	}

loop:
	for {
		// Peek the next token
		opToken := tokens[*index]

		switch opToken.Type {
		// Break if there are no tokens left
		case lx.Eof:
			break loop
		// Handle operator, Pratt Parsing
		case lx.Op:
			op, ok := opToken.Value.(lx.Operator)
			if !ok {
				panic(fmt.Sprintf("Bad token: %v\n", opToken))
			}

			// Check binding power
			lBp, rBp := op.BindingPower()
			if lBp < minBp {
				break loop
			}

			// We consume the op token now that everyting is in order
			*index++

			// Parse right expression
			rightNode := ParseExpression(tokens, index, rBp)

			leftNode = &Node{
				Type: Expression,
				Value: NodeExpression{
					Op:    string(op),
					Left:  leftNode,
					Right: rightNode,
				},
			}
		default:
			panic(fmt.Sprintf("Invalid op token: %v\n", opToken))
		}
	}

	return leftNode
}

func Parse(tokens []lx.Token) *Node {
	return ParseExpression(tokens, new(int), 0)
}
