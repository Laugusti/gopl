package eval

import "fmt"
import "math"
import "strings"

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

// A min computes the minimum value of its operands
type min struct {
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	// Eval returns the value of this Expr int the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
	// Pretty prints expression
	String() string
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() string {
	return string(v)
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	x := u.x.String()
	if strings.ContainsAny(x, "=-*/") {
		x = "(" + x + ")"
	}
	return fmt.Sprintf("%s%s", string(u.op), x)
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) String() string {
	x := b.x.String()
	if strings.ContainsAny(x, "=-*/") {
		x = "(" + x + ")"
	}
	y := b.y.String()
	if strings.ContainsAny(y, "=-*/") {
		y = "(" + y + ")"
	}

	return fmt.Sprintf("%s %s %s", x, string(b.op), y)
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	s := fmt.Sprintf("%s(", c.fn)
	for i, arg := range c.args {
		if i != 0 {
			s += fmt.Sprintf(", ")
		}
		s += fmt.Sprint(arg)
	}
	s += fmt.Sprint(")")
	return s
}

func (m min) Eval(env Env) float64 {
	min := m.args[0].Eval(env)
	for _, arg := range m.args {
		if val := arg.Eval(env); val < min {
			min = val
		}
	}
	return min
}

func (m min) Check(vars map[Var]bool) error {
	if len(m.args) < 1 {
		return fmt.Errorf("call to min requires at least 1 argument")
	}
	for _, arg := range m.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (m min) String() string {
	s := fmt.Sprint("min(")
	for i, arg := range m.args {
		if i != 0 {
			s += fmt.Sprintf(", ")
		}
		s += fmt.Sprint(arg)
	}
	s += fmt.Sprint(")")
	return s
}
