package parser_test

import (
	"rhodium/src/lexer"
	"rhodium/src/parser"
	"testing"
)

type TestCase struct {
	Tokens []lexer.Token
}

func TestPrattParse(t *testing.T) {
	testCases := []TestCase{
		{
			Tokens: []lexer.Token{
				{Type: lexer.Num, Value: 12},
				{Type: lexer.Op, Value: lexer.Plus},
				{Type: lexer.Num, Value: 5},
				{Type: lexer.Eof, Value: nil},
			},
		},
	}

	for _, testCase := range testCases {
		tree, _ := parser.PrattParse(testCase.Tokens, 0, 0)
		t.Logf("%#v\n", tree)
	}
}
