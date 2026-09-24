package lexer

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Tokenize(r io.Reader) []Token {
	// Buffered reader for peek method
	rd := bufio.NewReader(r)

	var tokens []Token

	for {
		b, err := rd.ReadByte()
		if err != nil {
			tokens = append(tokens, Token{Eof, nil})
			break
		}

		// Skip spaces and tabs
		if b == ' ' || b == '\t' {
			continue
		}

		// Parsing numbers
		if isDigit(b) {
			var numStr strings.Builder

			// Append any digits that we find
			for isDigit(b) {
				numStr.WriteByte(b)

				b, err = rd.ReadByte()
				if err != nil {
					break
				}
			}

			// TODO: Add logic to process floats

			// Parse the number and append the token
			num, err := strconv.Atoi(numStr.String())
			if err != nil {
				panic(err)
			}

			tokens = append(tokens, Token{Num, num})
		}

		// Parsing operators
		if isSymbol(b) {
			// So far we have operators that only have one symbol so this is fine
			tokens = append(tokens, Token{Op, Operator(b)})
		}
	}

	return tokens
}
