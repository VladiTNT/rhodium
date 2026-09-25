package lx_test

import (
	"rhodium/src/lx"
	"slices"
	"strings"
	"testing"
)

func TestValues(t *testing.T) {
	type TestCase struct {
		Input    string
		Expected []lx.Token
	}

	testCases := []TestCase{
		{
			Input: "1 2 3 44 5554432",
			Expected: []lx.Token{
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "1"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "2"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "3"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "44"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "5554432"}},
				{Type: lx.Eof, Value: nil},
			},
		},
		{
			Input: "1.2   22.6  3.1415   928173.213214 67.69",
			Expected: []lx.Token{
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "1.2"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "22.6"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "3.1415"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "928173.213214"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "67.69"}},
				{Type: lx.Eof, Value: nil},
			},
		},
		{
			Input: " 'vlad'  '1vv532'   '11 23ddrt' 'FnIJHjihbIUi7yYUb' ",
			Expected: []lx.Token{
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "vlad"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "1vv532"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "11 23ddrt"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "FnIJHjihbIUi7yYUb"}},
				{Type: lx.Eof, Value: nil},
			},
		},
		{
			Input: "22 'iFg7BHe22'   41 22.33 '132' 'u4TiE1as' 1.222222222   4123 'WDgaster' ",
			Expected: []lx.Token{
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "22"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "iFg7BHe22"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "41"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "22.33"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "132"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "u4TiE1as"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "1.222222222"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "4123"}},
				{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "WDgaster"}},
				{Type: lx.Eof, Value: nil},
			},
		},
	}

	for i, testCase := range testCases {
		tokens := lx.Tokenize(strings.NewReader(testCase.Input))

		// If the slices are the same then we pass
		if slices.Equal(testCase.Expected, tokens) {
			t.Logf("Test case %d passed!\n", i+1)
		} else {
			t.Errorf("Test case %d failed!\n", i+1)

			// If the lengths aren't the same then we point it out
			if len(testCase.Expected) != len(tokens) {
				t.Logf("Token slices don't have the same length, expected %d, got %d!\n",
					len(testCase.Expected), len(tokens))
			} else {
				// Log the token slices token for token
				for j := range len(testCase.Expected) {
					t.Logf("E %v G %v\n", testCase.Expected[j], tokens[j])
				}
			}
		}
	}
}
