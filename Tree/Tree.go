package Tree

import "fmt"

type TNode struct {
	Value    rune
	Children []*TNode
}

// NewTree creates an empty Tree
func NewTree() *TNode {
	return &TNode{0, []*TNode{}}
}

// IsEmpty checks whether a given tree is empty or not.
func (t *TNode) IsEmpty() bool {
	return (*t).Value == 0 && len((*t).Children) == 0
}

// NewNode creates a new node and returns its address.
func NewNode(val rune) *TNode {
	var tmp *TNode = &TNode{Value: val, Children: []*TNode{}}
	return tmp
}

// AddChild Adds new node as child of t.
func (t *TNode) AddChild(node *TNode) {
	t.Children = append(t.Children, node)
}

// String Display output the values present in the tree in out.
func String(node *TNode, out *string) {
	*out = (*out) + fmt.Sprintf("%c ", (*node).Value)
	for _, n := range (*node).Children {
		String(n, out)
	}
}
