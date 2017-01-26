package eval

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"min(0)", Env{}, "0"},
		{"min(A, B, C)", Env{"A": 1, "B": 2, "C": 3}, "1"},
		{"min(-2, 5/9, 3 * -2)", Env{}, "-6"},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.want {
			t.Errorf("%s.Eval() with %v = %q, want %q",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"min(0)", "call to min requires at least 1 argument"},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		err = expr.Check(map[Var]bool{})
		if err != nil && err.Error() != test.want {
			t.Errorf("%s.Check() = %q, want %q",
				test.expr, err, test.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"min(0)", Env{}, "min(0)"},
		{"min(A,B,C)", Env{"A": 1, "B": 2, "C": 3}, "min(A, B, C)"},
		{"min(-2, 5/9, 3 * -2)", Env{}, "min(-2, 5 / 9, 3 * (-2))"},
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
