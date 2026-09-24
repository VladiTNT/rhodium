package lexer

type Operator string

const (
	Plus  Operator = "+"
	Minus Operator = "-"

	Star     Operator = "*"
	Division Operator = "/"
	Modulus  Operator = "%"
)

func (o Operator) BindingPower() (float64, float64) {
	switch o {
	case Plus, Minus:
		return 1, 1.1
	case Star, Division, Modulus:
		return 2, 2.1
	default:
		panic("Invalid operator: " + o)
	}
}
