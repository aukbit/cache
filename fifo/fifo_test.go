package fifo

import "testing"

func TestIsEmpty(t *testing.T) {
	q := New(1)
	if !q.IsEmpty() {
		t.Fatal("It should be empty")
	}
}

func TestEnqueue1(t *testing.T) {
	q := New(3)
	a := &Item{
		Data: "A",
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
}

func TestEnqueue2(t *testing.T) {
	q := New(3)
	a := &Item{
		Data: "A",
	}
	b := &Item{
		Data: "B",
	}
	q.Enqueue(a)
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
}

func TestEnqueue3(t *testing.T) {
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
	q.Enqueue(a)
	q.Enqueue(b)
	q.Enqueue(c)
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	if q.first.Data.(string) != "C" {
		t.Fatal("First Item should be C")
	}
	if q.last.Data.(string) != "A" {
		t.Fatal("Last Item should be A")
	}
}

func TestDequeue1(t *testing.T) {
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
	if err := q.Enqueue(a); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue(b); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue(c); err != nil {
		t.Fatal(err)
	}
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	i := q.Dequeue()
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.Data.(string) != "C" {
		t.Fatal("First Item should be C")
	}
	if q.last.Data.(string) != "B" {
		t.Fatal("Last Item should be B")
	}
	if i.Data != "A" {
		t.Fatal("Dequeued Item should be A")
	}
	i = q.Dequeue()
	if i.Data != "B" {
		t.Fatal("Dequeued Item should be B")
	}
	i = q.Dequeue()
	if i.Data != "C" {
		t.Fatal("Dequeued Item should be C")
	}
	if !q.IsEmpty() {
		t.Fatal("Queue should be empty")
	}
}
