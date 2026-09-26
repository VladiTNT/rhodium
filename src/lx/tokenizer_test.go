package lx_test

import (
	"rhodium/src/lx"
	"slices"
	"strings"
	"testing"
)

type TestCase struct {
	Input    string
	Expected []lx.Token
}

var SubTests = []TestCase{
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
	{
		Input: "1 + 2 + 3",
		Expected: []lx.Token{
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "1"}},
			{Type: lx.Op, Value: lx.Plus},
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "2"}},
			{Type: lx.Op, Value: lx.Plus},
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "3"}},
			{Type: lx.Eof, Value: nil},
		},
	},
	{
		Input: "22 / 3 * 12 >= 9.53",
		Expected: []lx.Token{
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "22"}},
			{Type: lx.Op, Value: lx.Slash},
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "3"}},
			{Type: lx.Op, Value: lx.Star},
			{Type: lx.Atom, Value: lx.Value{Type: lx.Integer, Value: "12"}},
			{Type: lx.Op, Value: lx.GreaterThanEqual},
			{Type: lx.Atom, Value: lx.Value{Type: lx.Float, Value: "9.53"}},
			{Type: lx.Eof, Value: nil},
		},
	},
	{
		Input: " 'sS9kksd' + 'kl0tVzP' == 'sudo pacman' + '-Sybau' ",
		Expected: []lx.Token{
			{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "sS9kksd"}},
			{Type: lx.Op, Value: lx.Plus},
			{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "kl0tVzP"}},
			{Type: lx.Op, Value: lx.EqualEqual},
			{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "sudo pacman"}},
			{Type: lx.Op, Value: lx.Plus},
			{Type: lx.Atom, Value: lx.Value{Type: lx.String, Value: "-Sybau"}},
			{Type: lx.Eof, Value: nil},
		},
	},
}

func TestTokenizer(t *testing.T) {
	for i, subTest := range SubTests {
		tokens := lx.Tokenize(strings.NewReader(subTest.Input))

		if slices.Equal(subTest.Expected, tokens) {
			t.Logf("Sub test %d passed!\n", i+1)
		} else {
			t.Errorf("Sub test %d failed!\n", i+1)
			// Expected tokens
			for j, tk := range subTest.Expected {
				t.Logf("Expected token %d: %v\n", j, tk)
			}
			// Lexer output
			for j, tk := range tokens {
				t.Logf("Got token %d: %v\n", j, tk)
			}
		}
	}
}
