package lifo

import (
	"testing"
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
	if s.first.item != "A" {
		t.Fatalf("First item should be A not %v", s.first.item)
	}
	if s.last.item != "A" {
		t.Fatalf("Last item should be A not %v", s.last.item)
	}
	if err := s.Push("B"); err != nil {
		t.Fatal(err)
	}
	if s.first.item != "B" {
		t.Fatalf("First item should be B not %v", s.first.item)
	}
	if s.last.item != "A" {
		t.Fatalf("Last item should be A not %v", s.last.item)
	}
	if s.last.next != s.first {
		t.Fatalf("Last item next should be B not %v", s.first)
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
