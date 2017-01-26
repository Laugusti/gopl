package eval

import (
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A/pi)", Env{"A": 87616, "pi": math.Pi}, "sqrt(A / pi)"},
		{"pow(x,3)+pow(y,3)", Env{"x": 12, "y": 1}, "pow(x, 3) + pow(y, 3)"},
		{"5/9*(F-32)", Env{"F": 32}, "(5 / 9) * (F - 32)"},
		{"-(5/9)", Env{}, "-(5 / 9)"},
	}

	for _, test := range tests {
		expr1, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		if expr1.String() != test.want {
			t.Errorf("%v.String() in %v = %q, want %q\n",
				test.expr, test.env, expr1.String(), test.want)
		}
		expr2, err := Parse(expr1.String())
		if err != nil {
			t.Errorf("failed to parse: %s\n", err)
			continue
		}
		result1 := expr1.Eval(test.env)
		result2 := expr2.Eval(test.env)
		if result1 != result2 {
			t.Errorf("expr1.Eval() = %q, expr2,Eval() = %q\n",
				result1, result2)
		}
	}
}
