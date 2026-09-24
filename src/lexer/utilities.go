package lexer

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/' || b == '%'
}
