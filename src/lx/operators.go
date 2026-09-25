package lx

type Operator string

const (
	// Assign operator
	Equal Operator = "="

	// Level 1 arithmetic operators
	Plus  Operator = "+"
	Minus Operator = "-"

	// Level 2 arithmetic operators
	Star       Operator = "*"
	Slash      Operator = "/"
	Percentage Operator = "%"

	// Comparison operator
	LessThan         Operator = "<"
	GreaterThan      Operator = ">"
	EqualEqual       Operator = "=="
	LessThanEqual    Operator = "<="
	GreaterThanEqual Operator = ">="
)

// Binding power for each operator
func (o Operator) BindingPower() (float64, float64) {
	switch o {
	case Plus, Minus:
		return 1, 1.1
	case Star, Slash, Percentage:
		return 2, 2.2
	}
	panic("Operator with undefined binding power: " + o)
}
