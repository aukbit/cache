package fifo

import "testing"

func TestIsEmpty(t *testing.T) {
	q := New(1)
	if !q.IsEmpty() {
		t.Fatal("It should be empty")
	}
}

func TestQueue(t *testing.T) {
	q := New(3)
	a := &Item{
		Data: "A",
	}
	b := &Item{
		Data: "B",
	}
	c := &Item{
		Data: "C",
	}
	d := &Item{
		Data: "D",
	}
	// Add A
	q.Enqueue(a)
	if q.IsEmpty() {
		t.Fatal("It should not be empty")
	}
	if q.Size() != 1 {
		t.Fatal("It should be size 1")
	}
	if q.first.Data.(string) != "A" {
		t.Fatal("First Item should be A")
	}
	if q.last.Data.(string) != "A" {
		t.Fatal("Last Item should be A")
	}
	// Add B
	q.Enqueue(b)
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.Data.(string) != "B" {
		t.Fatal("First Item should be B")
	}
	if q.last.Data.(string) != "A" {
		t.Fatal("Last Item should be A")
	}
	// Add C
	q.Enqueue(c)
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	q.Enqueue(d)
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	if q.first.Data.(string) != "D" {
		t.Fatal("First Item should be D")
	}
	if q.last.Data.(string) != "B" {
		t.Fatal("Last Item should be B")
	}
	// Remove
	q.Dequeue()
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.Data.(string) != "D" {
		t.Fatal("First Item should be D")
	}
	if q.last.Data.(string) != "C" {
		t.Fatal("Last Item should be C")
	}
	q.Dequeue()
	if q.Size() != 1 {
		t.Fatal("It should be size 1")
	}
	if q.first.Data.(string) != "D" {
		t.Fatal("First Item should be D")
	}
	if q.last.Data.(string) != "D" {
		t.Fatal("Last Item should be D")
	}
	// Remove Last
	q.Dequeue()
	if q.Size() != 0 {
		t.Fatal("It should be size 0")
	}
	if q.first != nil {
		t.Fatal("First Item should be nil")
	}
	if q.last != nil {
		t.Fatal("Last Item should be nil")
	}
}
