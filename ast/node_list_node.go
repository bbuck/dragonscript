package ast

import (
	"github.com/bbuck/dragonscript"
	"github.com/bbuck/dragonscript/memory"
)

// NodeList is a special kind of node that represents a list of nodes. When
// a node list is evalutated in calls Eval on each node that it contains.
type NodeList []Node

// Eval will call Eval for every node inside of this NodeList passing in a new
// memory value.
func (nl NodeList) Eval(m memory.M) dragonscript.Value {
	var ret dragonscript.Value
	for _, n := range nl {
		ret = n.Eval(m)
	}

	return ret
}
