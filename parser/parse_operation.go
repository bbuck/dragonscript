package parser

import "github.com/bbuck/dragonscript/ast"

func (p *Parser) parseOperation(right ast.Node) (ast.Node, error) {
	op := p.nextToken()
	exp, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	return ast.NewOperatorNode(op, right, exp)
}
