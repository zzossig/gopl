// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) interface{} {
	return env[v]
}

func (l literal) Eval(_ Env) interface{} {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) interface{} {
	switch u.op {
	case '+':
		return u.x.Eval(env).(float64)
	case '-':
		return -1 * u.x.Eval(env).(float64)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) interface{} {
	switch b.op {
	case '+':
		return b.x.Eval(env).(float64) + b.y.Eval(env).(float64)
	case '-':
		return b.x.Eval(env).(float64) - b.y.Eval(env).(float64)
	case '*':
		return b.x.Eval(env).(float64) * b.y.Eval(env).(float64)
	case '/':
		return b.x.Eval(env).(float64) / b.y.Eval(env).(float64)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) interface{} {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env).(float64), c.args[1].Eval(env).(float64))
	case "sin":
		return math.Sin(c.args[0].Eval(env).(float64))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env).(float64))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Eval2

func (b boolean) Eval(env Env) interface{} {
	return b
}

func (c compare) Eval(env Env) interface{} {
	switch c.op {
	case '>':
		return c.x.Eval(env).(float64) > c.y.Eval(env).(float64)
	case '<':
		return c.x.Eval(env).(float64) < c.y.Eval(env).(float64)
	case '=':
		return c.x.Eval(env).(float64) == c.y.Eval(env).(float64)
	}
	panic(fmt.Sprintf("unsupported compare operator: %q", c.op))
}
