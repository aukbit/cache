package lru

import (
	"testing"

	"github.com/aukbit/cache"
)

func TestIsEmpty(t *testing.T) {
	c := New(1)
	if !c.IsEmpty() {
		t.Fatal("Cache should be empty")
	}
}

func TestSet(t *testing.T) {
	c := New(1)
	c.set("A", 1)
	if c.first.key != "A" {
		t.Fatalf("First Key should be A got %v", c.first.key)
	}
	if c.first.item != 1 {
		t.Fatalf("First Data should be 1 got %v", c.first.item)
	}
	if c.last.key != "A" {
		t.Fatalf("Last Key should be A got %v", c.last.key)
	}
	if c.last.item != 1 {
		t.Fatalf("Last Data should be 1 got %v", c.last.item)
	}
	if c.Size() != 1 {
		t.Fatalf("Size should be 1 got %v", c.Size())
	}
	if c.IsFull() != true {
		t.Fatalf("IsFull should be true got %v", c.IsFull())
	}
}

func TestSet2(t *testing.T) {
	c := New(2)
	c.set("A", 1)
	c.set("B", 2)
	if c.first.key != "B" {
		t.Fatalf("First Key should be B got %v", c.first.key)
	}
	if c.first.item != 2 {
		t.Fatalf("First Data should be 2 got %v", c.first.item)
	}
	if c.last.key != "A" {
		t.Fatalf("Last Key should be A got %v", c.last.key)
	}
	if c.last.item != 1 {
		t.Fatalf("Last Data should be 1 got %v", c.last.item)
	}
	if c.Size() != 2 {
		t.Fatalf("Size should be 2 got %v", c.Size())
	}
	if c.IsFull() != true {
		t.Fatalf("IsFull should be true got %v", c.IsFull())
	}
}

func TestSet3(t *testing.T) {
	c := New(2)
	// A
	c.set("A", 1)
	// B | A
	c.set("B", 2)
	// C | B
	c.set("C", 3)
	if c.first.key != "C" {
		t.Fatalf("First Key should be C got %v", c.first.key)
	}
	if c.first.item != 3 {
		t.Fatalf("First Data should be 3 got %v", c.first.item)
	}
	if c.last.key != "B" {
		t.Fatalf("Last Key should be B got %v", c.last.key)
	}
	if c.last.item != 2 {
		t.Fatalf("Last Data should be 2 got %v", c.last.item)
	}
	if c.Size() != 2 {
		t.Fatalf("Size should be 2 got %v", c.Size())
	}
	if c.IsFull() != true {
		t.Fatalf("IsFull should be true got %v", c.IsFull())
	}
}

func TestAccess(t *testing.T) {
	c := New(3)
	c.Access("A", 1)
	if c.first.item != 1 {
		t.Fatalf("first should be 1 got %v", c.first.item)
	}
	if c.last.item != 1 {
		t.Fatalf("last should be 1 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre != nil {
		t.Fatalf("pre should be nil got %v", c.first.pre)
	}
	// B | A
	c.Access("B", 2)
	if c.first.item != 2 {
		t.Fatalf("first should be 2 got %v", c.first.item)
	}
	if c.last.item != 1 {
		t.Fatalf("last should be 1 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 1 {
		t.Fatalf("pre should be 1 got %v", c.first.pre.item)
	}
	if c.last.next.item != 2 {
		t.Fatalf("next should be 2 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
	// C | B | A
	c.Access("C", 3)
	if c.first.item != 3 {
		t.Fatalf("first should be 3 got %v", c.first.item)
	}
	if c.last.item != 1 {
		t.Fatalf("last should be 1 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 2 {
		t.Fatalf("pre should be *b got %v", c.first.pre.item)
	}
	if c.last.next.item != 2 {
		t.Fatalf("next should be 2 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
	// B | C | A
	c.Access("B", 2)
	if c.first.item != 2 {
		t.Fatalf("first should be 2 got %v", c.first.item)
	}
	if c.last.item != 1 {
		t.Fatalf("last should be 1 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 3 {
		t.Fatalf("pre should be 3 got %v", c.first.pre.item)
	}
	if c.last.next.item != 3 {
		t.Fatalf("next should be 3 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
	// D | B | C
	c.Access("D", 4)
	if c.first.item != 4 {
		t.Fatalf("first should be 4 got %v", c.first.item)
	}
	if c.last.item != 3 {
		t.Fatalf("last should be 1 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 2 {
		t.Fatalf("pre should be 2 got %v", c.first.pre.item)
	}
	if c.last.next.item != 2 {
		t.Fatalf("next should be 2 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
	// C | D | B
	c.Access("C", 3)
	if c.first.item != 3 {
		t.Fatalf("first should be 3 got %v", c.first.item)
	}
	if c.last.item != 2 {
		t.Fatalf("last should be 2 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 4 {
		t.Fatalf("pre should be 4 got %v", c.first.pre.item)
	}
	if c.last.next.item != 4 {
		t.Fatalf("next should be 4 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
	// C | D | B
	c.Access("C", 3)
	if c.first.item != 3 {
		t.Fatalf("first should be 3 got %v", c.first.item)
	}
	if c.last.item != 2 {
		t.Fatalf("last should be 2 got %v", c.last.item)
	}
	if c.first.next != nil {
		t.Fatalf("next should be nil got %v", c.first.next)
	}
	if c.first.pre.item != 4 {
		t.Fatalf("pre should be 4 got %v", c.first.pre.item)
	}
	if c.last.next.item != 4 {
		t.Fatalf("next should be 4 got %v", c.last.next.item)
	}
	if c.last.pre != nil {
		t.Fatalf("pre should be nil got %v", c.last.pre)
	}
}

func TestIterator(t *testing.T) {
	b := New(2)
	b.Access("A", 1)
	b.Access("B", 2)
	i := b.Iterator()
	if !i.HasNext() {
		t.Fatal("Cache iterator should have next item")
	}
	c, err := i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != 1 {
		t.Fatalf("Iterator item should be 1 got %v", c)
	}
	if !i.HasNext() {
		t.Fatal("Cache iterator should have next item")
	}
	c, err = i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != 2 {
		t.Fatalf("Iterator item should be 2 got %v", c)
	}
	if i.HasNext() {
		t.Fatal("Should not be any more items")
	}
	if i.Remove() != cache.ErrUnsupportedOperation {
		t.Fatal("Remove operation should not be implemented")
	}
}
