package lifo

import (
	"testing"

	"github.com/aukbit/cache"
)

func TestIsEmpty(t *testing.T) {
	s := New(1)
	if !s.IsEmpty() {
		t.Fatal("Stack should be empty")
	}
}

func TestPush1(t *testing.T) {
	s := New(1)
	if err := s.Push("A"); err != nil {
		t.Fatal(err)
	}
	if s.Size() != 1 {
		t.Fatalf("Stack size should be 1 not %v", s.Size())
	}
}

func TestPush2(t *testing.T) {
	s := New(2)
	if err := s.Push("A"); err != nil {
		t.Fatal(err)
	}
	if s.last.item != "A" {
		t.Fatalf("Last item should be A not %v", s.last.item)
	}
	if err := s.Push("B"); err != nil {
		t.Fatal(err)
	}
	if s.last.item != "B" {
		t.Fatalf("Last item should be B not %v", s.last.item)
	}
	if s.last.next.item != "A" {
		t.Fatalf("Last item next should be A not %v", s.last.next.item)
	}
	if s.Size() != 2 {
		t.Fatalf("Stack size should be 2 not %v", s.Size())
	}
}

func TestFull(t *testing.T) {
	s := New(2)
	if err := s.Push("A"); err != nil {
		t.Fatal(err)
	}
	if err := s.Push("B"); err != nil {
		t.Fatal(err)
	}
	if err := s.Push("C"); err != ErrStackIsFull {
		t.Fatal(err)
	}
	if s.Size() != 2 {
		t.Fatalf("Stack size should be 2 not %v", s.Size())
	}
}

func TestPop(t *testing.T) {
	q := New(3)
	if err := q.Push("A"); err != nil {
		t.Fatal(err)
	}
	if err := q.Push("B"); err != nil {
		t.Fatal(err)
	}
	if err := q.Push("C"); err != nil {
		t.Fatal(err)
	}
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	i := q.Pop()
	if i != "C" {
		t.Fatalf("Pop item should be C got %v", i)
	}
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.last.item != "B" {
		t.Fatalf("Last item should be B got %v", q.last.item)
	}
	i = q.Pop()
	if i != "B" {
		t.Fatalf("Pop item should be B got %v", i)
	}
	i = q.Pop()
	if i != "A" {
		t.Fatalf("Pop item should be A got %v", i)
	}
	if !q.IsEmpty() {
		t.Fatal("Stack should be empty")
	}
}

func TestIterator(t *testing.T) {
	b := New(2)
	if err := b.Push("A"); err != nil {
		t.Fatal(err)
	}
	if err := b.Push("B"); err != nil {
		t.Fatal(err)
	}
	i := b.Iterator()
	if !i.HasNext() {
		t.Fatal("Stack iterator should have next item")
	}
	c, err := i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != "B" {
		t.Fatalf("Iterator item should be B got %v", c)
	}
	if !i.HasNext() {
		t.Fatal("Stack iterator should have next item")
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
