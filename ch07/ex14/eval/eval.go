package eval

import (
	"fmt"
	"math"
)

type Var string
type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

type list struct {
	fn   string
	args []Expr
}

type ternery struct {
	op1, op2 rune
	x, y, z  Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	String() string
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) String() string {
	return string(v)
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
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

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
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
func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	case "max":
		return math.Max(c.args[0].Eval(env), c.args[1].Eval(env))
	case "min":
		return math.Min(c.args[0].Eval(env), c.args[1].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (l list) Eval(env Env) float64 {
	switch l.fn {
	case "min":
		min := l.args[0].Eval(env)
		for _, arg := range l.args[1:] {
			val := arg.Eval(env)
			if min > val {
				min = val
			}
		}
		return min
	case "max":
		max := l.args[0].Eval(env)
		for _, arg := range l.args[1:] {
			val := arg.Eval(env)
			if max < val {
				max = val
			}
		}
		return max
	}
	panic(fmt.Sprintf("unsupported list call: %s", l.fn))
}

func (l list) String() string {
	args := ""
	for i, arg := range l.args {
		if i < len(l.args)-1 {
			args += fmt.Sprintf("%s, ", arg)
		} else {
			args += fmt.Sprintf("%s", arg)
		}
	}
	return fmt.Sprintf("%s<%s>", l.fn, args)
}

func (t ternery) Eval(env Env) float64 {
	if t.op1 == '?' && t.op2 == ':' {
		if t.x.Eval(env) != 0 {
			return t.y.Eval(env)
		}
		return t.z.Eval(env)
	}
	panic(fmt.Sprintf("unsupported ternery operator: %q, %q", t.op1, t.op2))
}

func (t ternery) String() string {
	return fmt.Sprintf("(%s %c %s %c %s)", t.x, t.op1, t.y, t.op2, t.z)
}

func (c call) String() string {
	args := ""
	for i, arg := range c.args {
		if i < len(c.args)-1 {
			args += fmt.Sprintf("%s, ", arg)
		} else {
			args += fmt.Sprintf("%s", arg)
		}
	}
	return fmt.Sprintf("%s(%s)", c.fn, args)
}
