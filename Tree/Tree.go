package Tree

import "fmt"

type TNode struct {
	Value    rune
	Children []*TNode
}

func NewNode(val rune) *TNode {
	var tmp *TNode = &TNode{Value: val, Children: []*TNode{}}
	return tmp
}

// AddChild Adds new node as child of t.
func (t *TNode) AddChild(node *TNode) {
	t.Children = append(t.Children, node)
}

// Display output the values present in the tree in out.
func String(node *TNode, out *string) {
	*out = (*out) + fmt.Sprintf("%c ", (*node).Value)
	for _, n := range (*node).Children {
		String(n, out)
	}
}
