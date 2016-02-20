package ast

import (
	"strconv"

	"github.com/bbuck/dragonscript"
	"github.com/bbuck/dragonscript/lexer"
	"github.com/bbuck/dragonscript/memory"
)

// IntNode is a DragonScript value that represents an integer in go.
type IntNode struct {
	token *lexer.Token
	value int64
}

// NewIntNode creates an new integer node, pre-parsing the string value as an integer.
func NewIntNode(t *lexer.Token) (*IntNode, error) {
	i := &IntNode{token: t}
	num, err := strconv.ParseInt(purgeNumericStrings(t.Value), 0, 64)
	if err != nil {
		return nil, err
	}
	i.value = num

	return i, nil
}

// Eval returns the DragonScript integer value this node represents
func (i *IntNode) Eval(m memory.M) dragonscript.Value {
	return dragonscript.Integer(i.value)
}
