package main

import (
	"flag"
	"fmt"

	"github.com/bbuck/dragonscript/lexer"
	"github.com/bbuck/dragonscript/parser"
)

var test = flag.String("test", "", "the string to test")

func main() {
	flag.Parse()

	tokens, err := lexer.Tokenize(*test)
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)
	p := parser.NewParser(tokens)
	tree, err := p.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", tree.Eval(nil))
}
