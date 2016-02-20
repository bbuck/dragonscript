package parser

import (
	"fmt"

	"github.com/bbuck/dragonscript/ast"
	"github.com/bbuck/dragonscript/lexer"
)

func createNumericNode(t *lexer.Token) (ast.Node, error) {
	switch t.Type {
	case lexer.TokenInt:
		return ast.NewIntNode(t)
	case lexer.TokenFloat:
		return ast.NewFloatNode(t)
	default:
		return nil, fmt.Errorf("tried to handle %s token as numeric", t.Type)
	}
}

func isTerminatorToken(t *lexer.Token) bool {
	return t.Type == lexer.TokenEOF || t.Type == lexer.TokenTerminator
}
