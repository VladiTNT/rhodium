package lexer

type Operator string

const (
	Plus  Operator = "+"
	Minus Operator = "-"

	Star     Operator = "*"
	Division Operator = "/"
	Modulus  Operator = "%"
)
