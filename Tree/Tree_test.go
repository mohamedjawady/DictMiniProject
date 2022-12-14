package Tree

import "testing"

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
