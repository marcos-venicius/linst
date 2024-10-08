package linst

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	tree := Create[int]()

	r := tree.IsEmpty()

	if !r {
		t.Fatal("Expected: true, Received: false")
	}

	tree.Add(1)

	r = tree.IsEmpty()

	if r {
		t.Fatal("Expected: false, Received: true")
	}
}

func TestTreeShouldBeCreatedWithNilValues(t *testing.T) {
	tree := Create[int]()

	if tree.root != nil {
		t.Fatal("tree was not created with nil root")
	}

	if tree.current != nil {
		t.Fatal("tree was not created with nil current")
	}
}

func TestRootAndCurrentShouldPointToSameObjectWhenAddFirstTime(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)

	if &*tree.root != &*tree.current {
		t.Fatalf(
			"root: %p, current: %p. They should be the same",
			&*tree.root,
			&*tree.current,
		)
	}
}

func TestRootShouldBeTheSameWhenAddingMoreThanOne(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)

	root := &*tree.root

	tree.Add(2)
	tree.Add(3)

	if &*tree.root != root {
		t.Fatalf(
			"old root: %p, new root: %p. They should be the same",
			root,
			&*tree.root,
		)
	}
}

func TestCurrentShouldBeLastItemAdded(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	if &*tree.current != tree.root.next.next {
		t.Fatalf(
			"current: %p, last item: %p. They should be the same",
			&*tree.current,
			&tree.root.next.next,
		)
	}
}

func TestAddMiddle(t *testing.T) {
	tree := Create[float64]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.Prev()

	tree.Add(1.5)

	tree.SelectRoot()

	v, _ := tree.Next()

	if v.Data != 1.5 {
		t.Fatalf("Expected: 1.5, Received: %f", v.Data)
	}

	if &*v.prev != tree.root {
		t.Fatalf("%p != %p", &*v.prev, tree.root)
	}

	if &*v.next != tree.root.next.next {
		t.Fatalf("%p != %p", &*v.prev, tree.root.next.next)
	}
}

func TestAddBeginning(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)

	tree.Prev()

	tree.Add(0)

	tree.SelectRoot()

	node := tree.Node()

	if node.Data != 0 {
		t.Fatalf("Expected: 0, Received: %d", node.Data)
	}

	if &*tree.root != node {
		t.Fatalf("%p != %p", &*tree.root, node)
	}
}

func TestCurrentShouldHaveNextNilWhenAddingLast(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(3)
	tree.Add(3)
	tree.Add(3)

	if tree.current.next != nil {
		t.Fatal("Next current should be nil")
	}
}

func TestRootPrevShouldBeNil(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(3)
	tree.Add(3)
	tree.Add(3)

	if tree.root.prev != nil {
		t.Fatal("Prev root should be nil")
	}
}

func TestShouldReturnPrevCorrectly(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	v, _ := tree.Prev()

	if v.Data != 2 {
		t.Fatalf("Expected: 2, Received: %d", v.Data)
	}

	v, _ = tree.Prev()

	if v.Data != 1 {
		t.Fatalf("Expected: 1, Received: %d", v.Data)
	}
}

func TestShouldReturnErrorWhenDoesNotHavePrevAnymore(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	_, err := tree.Prev()

	if err != nil {
		t.Fatal("Should not return error")
	}

	_, err = tree.Prev()

	if err != nil {
		t.Fatal("Should not return error")
	}

	_, err = tree.Prev()

	expectedError := "prev node not found"

	if err == nil {
		t.Fatal("Should return error")
	} else if err.Error() != expectedError {
		t.Fatalf("Expected: %v, Received: %v", expectedError, err.Error())
	}
}

func TestShouldReturnNextCorrectly(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.Prev()
	tree.Prev()
	tree.Prev()

	v, _ := tree.Next()

	if v.Data != 2 {
		t.Fatalf("Expected: 2, Received: %d", v.Data)
	}

	v, _ = tree.Next()

	if v.Data != 3 {
		t.Fatalf("Expected: 3, Received: %d", v.Data)
	}
}

func TestShouldReturnErrorWhenDoesNotHaveNextAnymore(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.Prev()

	_, err := tree.Next()

	if err != nil {
		t.Fatal("Should not return error")
	}

	_, err = tree.Next()

	expectedError := "next node not found"

	if err == nil {
		t.Fatal("Should return error")
	} else if err.Error() != expectedError {
		t.Fatalf("Expected: %v, Received: %v", expectedError, err.Error())
	}
}

func TestHasPrev(t *testing.T) {
	tree := Create[int]()

	r := tree.HasPrev()

	if r {
		t.Fatal("Expected: false, Received: true")
	}

	tree.Add(1)

	r = tree.HasPrev()

	if r {
		t.Fatal("Expected: false, Received: true")
	}

	tree.Add(2)

	r = tree.HasPrev()

	if !r {
		t.Fatal("Expected: true, Received: false")
	}
}

func TestHasNext(t *testing.T) {
	tree := Create[int]()

	r := tree.HasNext()

	if r {
		t.Fatal("Expected: false, Received: true")
	}

	tree.Add(1)

	r = tree.HasNext()

	if r {
		t.Fatal("Expected: false, Received: true")
	}

	tree.Add(2)

	tree.Prev()

	r = tree.HasNext()

	if !r {
		t.Fatal("Expected: true, Received: false")
	}
}

func TestNode(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	if tree.Node() != &*tree.current {
		t.Fatalf("Expected: %p, Received: %p", &*tree.current, tree.Node())
	}
}

func TestRoot(t *testing.T) {

	tree := Create[int]()

	root := tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	if tree.Root() != &*root {
		t.Fatal("invalid root")
	}
}

func TestSelectRoot(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.SelectRoot()

	node := tree.Node()

	if node == nil {
		t.Fatal("Node should not be nil")
	} else if node.Data != 1 {
		t.Fatalf("Expected: %d, Received: %d", 1, node.Data)
	}
}

func TestDeleteBeginning(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	s := tree.Add(3)

	tree.SelectRoot()

	tree.Delete()

	if tree.root.Data != 2 {
		t.Fatalf("Expected: 2, Received: %d", tree.root.Data)
	}

	if tree.root.prev != nil {
		t.Fatal("Root prev should always be nil")
	}

	if tree.root.next != s {
		t.Fatal("Invalid tree")
	}
}

func TestDeleteUnique(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)

	tree.Delete()

	if tree.root != nil {
		t.Fatal("root should be nil")
	}
}

func TestDeleteMiddle(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	v, _ := tree.Prev()

	if v.Data != 2 {
		t.Fatal("invalid prev")
	}

	tree.Delete()

	v = tree.Node()

	if v.Data != 3 {
		t.Fatal("Expect 3")
	}

	if v.prev.Data != 1 {
		t.Fatal("invalid prev node")
	}

	if v.next != nil {
		t.Fatal("invalid next node")
	}
}

func TestDeleteEnd(t *testing.T) {
	tree := Create[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	tree.Delete()

	v := tree.Node()

	if v.Data != 2 {
		t.Fatal("invalid selected node")
	}

	if v.next != nil {
		t.Fatal("invalid next node")
	}

	if v.prev.Data != 1 {
		t.Fatal("invalid prev node")
	}
}

func TestDeleteEmpty(t *testing.T) {
	tree := Create[int]()

	err := tree.Delete()

	if err == nil {
		t.Fatal("Expected error")
	} else if err.Error() != "Tree is empty" {
		t.Fatal("invalid error message")
	}
}
