package evaluator

import (
	"github.com/bin_lang/ast"
	"github.com/bin_lang/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
