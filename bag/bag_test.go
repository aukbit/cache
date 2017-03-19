package bag

import (
	"testing"

	"github.com/aukbit/cache"
)

func TestIsEmpty(t *testing.T) {
	b := New(1)
	if !b.IsEmpty() {
		t.Fatal("Bag should be empty")
	}
}

func TestAdd1(t *testing.T) {
	b := New(1)
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	if b.Size() != 1 {
		t.Fatalf("Bag size should be 1 not %v", b.Size())
	}
}

func TestAdd2(t *testing.T) {
	b := New(2)
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	if err := b.Add("B"); err != nil {
		t.Fatal(err)
	}
	if b.Size() != 2 {
		t.Fatalf("Bag size should be 2 not %v", b.Size())
	}
	if err := b.Add("C"); err == nil {
		t.Fatal("Bag should be full")
	}
}

func TestIterator(t *testing.T) {
	b := New(2)
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	i := b.Iterator()
	if !i.HasNext() {
		t.Fatal("Bag iterator should have first Item")
	}
	c, err := i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if b.first.item != c {
		t.Fatal("Bag iterator should have first Item")
	}
	if i.Remove() != cache.ErrUnsupportedOperation {
		t.Fatal("Remove operation should not be implemented")
	}
	_, err = i.Next()
	if err != cache.ErrNoSuchElement {
		t.Fatal("Should not be any more items")
	}
}
