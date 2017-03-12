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
	a := &Item{
		Data: "A",
	}
	if err := s.Push(a); err != nil {
		t.Fatal(err)
	}
	if s.Size() != 1 {
		t.Fatalf("Stack size should be 1 not %v", s.Size())
	}
}

func TestPush2(t *testing.T) {
	s := New(2)
	a := &Item{
		Data: "A",
	}
	b := &Item{
		Data: "B",
	}
	if err := s.Push(a); err != nil {
		t.Fatal(err)
	}
	if s.first.Data != "A" {
		t.Fatalf("First item should be A not %v", s.first.Data)
	}
	if s.last.Data != "A" {
		t.Fatalf("Last item should be A not %v", s.last.Data)
	}
	if err := s.Push(b); err != nil {
		t.Fatal(err)
	}
	if s.first.Data != "B" {
		t.Fatalf("First item should be B not %v", s.first.Data)
	}
	if s.last.Data != "A" {
		t.Fatalf("Last item should be A not %v", s.last.Data)
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
	a := &Item{
		Data: "A",
	}
	b := &Item{
		Data: "B",
	}
	c := &Item{
		Data: "C",
	}
	if err := s.Push(a); err != nil {
		t.Fatal(err)
	}
	if err := s.Push(b); err != nil {
		t.Fatal(err)
	}
	if err := s.Push(c); err != ErrStackIsFull {
		t.Fatal(err)
	}
	if s.Size() != 2 {
		t.Fatalf("Stack size should be 2 not %v", s.Size())
	}
}
