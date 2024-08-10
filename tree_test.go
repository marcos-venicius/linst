package main

import (
	"testing"
)

func TestTreeShouldBeCreatedWithNilValues(t *testing.T) {
	tree := CreateTree[int]()

	if tree.root != nil {
		t.Fatal("tree was not created with nil root")
	}

	if tree.current != nil {
		t.Fatal("tree was not created with nil current")
	}
}

func TestRootAndCurrentShouldPointToSameObjectWhenAddFirstTime(t *testing.T) {
	tree := CreateTree[int]()

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
	tree := CreateTree[int]()

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
	tree := CreateTree[int]()

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

func TestCurrentShouldHaveNextNilWhenAddingLast(t *testing.T) {
	tree := CreateTree[int]()

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
	tree := CreateTree[int]()

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
	tree := CreateTree[int]()

	tree.Add(1)
	tree.Add(2)
	tree.Add(3)

	v, _ := tree.Prev()

	if v.data != 2 {
		t.Fatalf("Expected: 2, Received: %d", v.data)
	}

	v, _ = tree.Prev()

	if v.data != 1 {
		t.Fatalf("Expected: 1, Received: %d", v.data)
	}
}

func TestShouldReturnErrorWhenDoesNotHavePrevAnymore(t *testing.T) {
	tree := CreateTree[int]()

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
