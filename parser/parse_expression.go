package parser

import (
	"fmt"

	"github.com/bbuck/dragonscript/ast"
	"github.com/bbuck/dragonscript/lexer"
)

func (p *Parser) parseExpression() (ast.Node, error) {
	token := p.nextToken()

	switch {
	case token == nil:
		return nil, nil
	case token.Type == lexer.TokenInt || token.Type == lexer.TokenFloat:
		next := p.peek(1)
		right, err := createNumericNode(token)
		if err != nil {
			return nil, err
		}

		switch {
		case next.Type == lexer.TokenOperator:
			return p.parseOperation(right)
		case isTerminatorToken(next):
			fallthrough
		default:
			p.nextToken()

			return right, err
		}
	case token.Type == lexer.TokenEOF:
		return nil, nil
	default:
		return nil, fmt.Errorf("Unexpected token %s found", token.Type)
	}
}
