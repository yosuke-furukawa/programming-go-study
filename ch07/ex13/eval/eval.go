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
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
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
