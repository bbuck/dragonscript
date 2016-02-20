package ast

import (
	"strconv"

	"github.com/bbuck/dragonscript"
	"github.com/bbuck/dragonscript/lexer"
	"github.com/bbuck/dragonscript/memory"
)

// FloatNode defines a node for representing a FloatNode value inside of DragonScript.
type FloatNode struct {
	token *lexer.Token
	value float64
}

// NewFloatNode creats a new float node and parses it's value for use later.
func NewFloatNode(t *lexer.Token) (*FloatNode, error) {
	f := &FloatNode{token: t}
	num, err := strconv.ParseFloat(purgeNumericStrings(t.Value), 64)
	if err != nil {
		return nil, err
	}
	f.value = num

	return f, nil
}

// Eval returns the float that is associated with this node and usable by
// Go.
func (f *FloatNode) Eval(m memory.M) dragonscript.Value {
	return dragonscript.Float(f.value)
}
