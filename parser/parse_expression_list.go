package parser

import "github.com/bbuck/dragonscript/ast"

func (p *Parser) parseExpressionList() (ast.Node, error) {
	list := make(ast.NodeList, 0)
	var (
		exp ast.Node
		err error
	)
	for exp, err = p.parseExpression(); exp != nil && err == nil; exp, err = p.parseExpression() {
		list = append(list, exp)
	}

	return list, err
}
