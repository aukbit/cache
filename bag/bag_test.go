package bag

import (
	"testing"

	"github.com/aukbit/cache"
)

func TestIsEmpty(t *testing.T) {
	b := New()
	if !b.IsEmpty() {
		t.Fatal("Bag should be empty")
	}
}

func TestAdd1(t *testing.T) {
	b := New()
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	if b.Size() != 1 {
		t.Fatalf("Bag size should be 1 not %v", b.Size())
	}
}

func TestAdd2(t *testing.T) {
	b := New()
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	if err := b.Add("B"); err != nil {
		t.Fatal(err)
	}
	if b.Size() != 2 {
		t.Fatalf("Bag size should be 2 not %v", b.Size())
	}
}

func TestIterator(t *testing.T) {
	b := New()
	if err := b.Add("A"); err != nil {
		t.Fatal(err)
	}
	if err := b.Add("B"); err != nil {
		t.Fatal(err)
	}
	i := b.Iterator()
	if !i.HasNext() {
		t.Fatal("Bag iterator should have next item")
	}
	c, err := i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != "B" {
		t.Fatalf("Iterator item should be B got %v", c)
	}
	if !i.HasNext() {
		t.Fatal("Bag iterator should have next item")
	}
	c, err = i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != "A" {
		t.Fatalf("Iterator item should be A got %v", c)
	}
	if i.HasNext() {
		t.Fatal("Should not be any more items")
	}
	if i.Remove() != cache.ErrUnsupportedOperation {
		t.Fatal("Remove operation should not be implemented")
	}
}
