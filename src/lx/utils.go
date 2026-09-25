package lx

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	switch b {
	case '=', '+', '-', '*', '/', '%', '<', '>':
		return true
	}
	return false
}

func isApostrophe(b byte) bool {
	return b == '\''
}
