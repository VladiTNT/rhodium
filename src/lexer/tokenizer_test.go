package lexer_test

import (
	"rhodium/src/lexer"
	"slices"
	"strings"
	"testing"
)

type TestCase struct {
	Expression string
	Tokens     []lexer.Token
}

func TestBasicExpression(t *testing.T) {
	testCases := []TestCase{
		{
			Expression: "123 + 22 / 44 * 55",
			Tokens: []lexer.Token{
				{lexer.Num, 123},
				{lexer.Op, lexer.Plus},
				{lexer.Num, 22},
				{lexer.Op, lexer.Division},
				{lexer.Num, 44},
				{lexer.Op, lexer.Star},
				{lexer.Num, 55},
				{lexer.Eof, nil},
			},
		},
	}

	for i, testCase := range testCases {
		tokensFromExpression := lexer.Tokenize(strings.NewReader(testCase.Expression))
		if !slices.Equal(testCase.Tokens, tokensFromExpression) {
			t.Errorf("Test case %d failed.\n", i+1)

			testTokensLen := len(testCase.Tokens)
			tokensFromExpressionLen := len(tokensFromExpression)

			if testTokensLen == tokensFromExpressionLen {
				for j := range testTokensLen {
					t.Logf("wanted %#v, got %#v\n", testCase.Tokens[j], tokensFromExpression[j])
				}
			} else {
				t.Errorf("Token slices don't even have the same length: wanted %d, got %d.\n", testTokensLen, tokensFromExpressionLen)
			}
		} else {
			t.Logf("Test case %d passed.\n", i+1)
		}
	}
}
