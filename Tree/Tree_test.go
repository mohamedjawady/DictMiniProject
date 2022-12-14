package Tree

import "testing"

func TestIsEmpty(t *testing.T) {
	type testEmpty struct {
		value    bool
		expected bool
	}
	tree1 := NewTree()
	tree2 := NewNode('a')
	var tests = []testEmpty{
		{tree1.IsEmpty(), true},
		{tree2.IsEmpty(), false},
	}
	for _, test := range tests {
		if test.value != test.expected {
			t.Errorf("Output %t not equal to expected %t", test.value, test.expected)
		}
	}
}
func TestNewNode(t *testing.T) {
	type newTest struct {
		value    rune
		expected string
	}
	var testNodes = []newTest{
		{'a', "a "},
		{'A', "A "},
	}
	for _, test := range testNodes {
		tree := NewNode(test.value)
		var output string
		String(tree, &output)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestAddChild(t *testing.T) {
	tree := NewNode('a')
	nodeB := NewNode('B')
	tree.AddChild(nodeB)
	var got string
	String(tree, &got)
	want := "a B "
	if got != want {
		t.Errorf("Output %q not equal to expected %q", got, want)
	}
}
