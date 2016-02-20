package ast

import (
	"fmt"
	"regexp"

	"github.com/bbuck/dragonscript"
	"github.com/bbuck/dragonscript/lexer"
	"github.com/bbuck/dragonscript/memory"
)

var validOperators = regexp.MustCompile(`==?|\*\*=?|\*=?|\+=?|/=?|-=?|!(?:=|~)?|~=|<<=?|<=?|>>=?|>=?|\^=?`)

type OperatorNode struct {
	token *lexer.Token
	right Node
	left  Node
}

func NewOperatorNode(t *lexer.Token, r Node, l Node) (*OperatorNode, error) {
	if !validOperators.MatchString(t.Value) {
		return nil, fmt.Errorf("operator %q is not valid", t.Value)
	}

	return &OperatorNode{
		token: t,
		right: r,
		left:  l,
	}, nil
}

func (o *OperatorNode) Eval(m memory.M) dragonscript.Value {
	right := o.right.Eval(m)
	left := o.left.Eval(m)

	// for now, only numbers
	isFloatOp := false
	_, rok := right.(dragonscript.Float)
	_, lok := left.(dragonscript.Float)
	if rok || lok {
		isFloatOp = true
	}

	if isFloatOp {
		return floatResult(o.token, right, left)
	}

	return intResult(o.token, right, left)
}

func floatFromValue(v dragonscript.Value) float64 {
	if f, ok := v.(dragonscript.Float); ok {
		return f.GoValue().(float64)
	}

	i := v.GoValue().(int64)

	return float64(i)
}

func intFromValue(v dragonscript.Value) int64 {
	if i, ok := v.(dragonscript.Integer); ok {
		return i.GoValue().(int64)
	}

	f := v.GoValue().(float64)

	return int64(f)
}

func floatResult(t *lexer.Token, r dragonscript.Value, l dragonscript.Value) dragonscript.Value {
	var (
		res float64
		rf  = floatFromValue(r)
		lf  = floatFromValue(l)
	)
	switch t.Value {
	case "+":
		res = rf + lf
	case "-":
		res = rf - lf
	case "*":
		res = rf * lf
	case "/":
		res = rf / lf
	}

	return dragonscript.Float(res)
}

func intResult(t *lexer.Token, r dragonscript.Value, l dragonscript.Value) dragonscript.Value {
	var (
		res int64
		ri  = intFromValue(r)
		li  = intFromValue(l)
	)
	switch t.Value {
	case "+":
		res = ri + li
	case "-":
		res = ri - li
	case "*":
		res = ri * li
	case "/":
		res = ri / li
	}

	return dragonscript.Integer(res)
}
