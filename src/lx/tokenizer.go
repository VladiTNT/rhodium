package lx

import (
	"bufio"
	"io"
	"strings"
)

func Tokenize(r io.Reader) []Token {
	rd := bufio.NewReader(r)

	var tokens []Token

	for {
		b, err := rd.ReadByte()
		if err != nil {
			break
		}

		// Skip spaces
		if b == ' ' || b == '\t' {
			continue
		}

		// Parse numbers
		if isDigit(b) {
			var numStr strings.Builder
			var currType ValueType = Integer

			// While we get digits
			for isDigit(b) {
				numStr.WriteByte(b)

				b, err = rd.ReadByte()
				if err != nil {
					break
				}
			}

			// Check if the byte that broke the thing is a dot because then
			// we are dealing with a float
			if b == '.' {
				// Make the value a float
				currType = Float

				// Write the dot to the value
				numStr.WriteByte(b)

				// Add the digits after the dot
				for {
					b, err = rd.ReadByte()
					if err != nil {
						break
					}

					if isDigit(b) {
						numStr.WriteByte(b)
					} else {
						break
					}
				}
			}

			// Add the integer or float that we just parsed
			tokens = append(tokens, Token{Atom, Value{currType, numStr.String()}})

			continue
		}

		// Parse operators
		if isSymbol(b) {
			var opStr strings.Builder

			// Add this first symbol
			opStr.WriteByte(b)

			// Peek the next byte for operators made out of 2 characters
			bb, err := rd.Peek(1)
			if err == nil {
				// If the first byte was a
				switch b {
				// '<=' '>=' '=='
				case '<', '>', '=':
					if bb[0] == '=' {
						opStr.WriteByte(bb[0])
					}
				}
			}

			// Add the operator
			tokens = append(tokens, Token{Op, Operator(opStr.String())})
			continue
		}

		// Parse string literals
		if isApostrophe(b) {
			var strStr strings.Builder

			b, err = rd.ReadByte()
			if err != nil {
				break
			}

			// Read all bytes until the next apostrophe
			for !isApostrophe(b) {
				strStr.WriteByte(b)

				b, err = rd.ReadByte()
				if err != nil {
					break
				}
			}

			// Add the string token
			tokens = append(tokens, Token{Atom, Value{String, strStr.String()}})

			// Clear the closing apostrophe
			_, err := rd.ReadByte()
			if err != nil {
				break
			}
		}
	}

	// The token list must end with an Eof token
	tokens = append(tokens, Token{Eof, nil})

	return tokens
}
