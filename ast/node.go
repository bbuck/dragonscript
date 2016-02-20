package ast

import (
	"github.com/bbuck/dragonscript"
	"github.com/bbuck/dragonscript/memory"
)

// A Node is an interface that defines an object that has an Eval function
type Node interface {
	Eval(memory.M) dragonscript.Value
}
