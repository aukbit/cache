package fifo

import (
	"testing"

	"github.com/aukbit/cache"
)

func TestIsEmpty(t *testing.T) {
	q := New(1)
	if !q.IsEmpty() {
		t.Fatal("It should be empty")
	}
}

func TestEnqueue1(t *testing.T) {
	q := New(3)
	q.Enqueue("A")
	if q.IsEmpty() {
		t.Fatal("It should not be empty")
	}
	if q.Size() != 1 {
		t.Fatal("It should be size 1")
	}
	if q.first.item != "A" {
		t.Fatal("First Item should be A")
	}
	if q.last.item != "A" {
		t.Fatal("Last Item should be A")
	}
}

func TestEnqueue2(t *testing.T) {
	q := New(3)
	q.Enqueue("A")
	q.Enqueue("B")
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.item != "B" {
		t.Fatal("First Item should be B")
	}
	if q.last.item != "A" {
		t.Fatal("Last Item should be A")
	}
	if q.last.next != q.first {
		t.Fatalf("Last item next should be B not %v", q.first)
	}
}

func TestEnqueue3(t *testing.T) {
	q := New(3)
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	if q.first.item != "C" {
		t.Fatal("First Item should be C")
	}
	if q.last.item != "A" {
		t.Fatal("Last Item should be A")
	}
}

func TestFull(t *testing.T) {
	q := New(2)
	if err := q.Enqueue("A"); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue("B"); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue("C"); err != ErrStackIsFull {
		t.Fatal(err)
	}
	if q.Size() != 2 {
		t.Fatalf("Queue size should be 2 not %v", q.Size())
	}
}

func TestDequeue1(t *testing.T) {
	q := New(3)
	if err := q.Enqueue("A"); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue("B"); err != nil {
		t.Fatal(err)
	}
	if err := q.Enqueue("C"); err != nil {
		t.Fatal(err)
	}
	if q.Size() != 3 {
		t.Fatal("It should be size 3")
	}
	i := q.Dequeue()
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.item != "C" {
		t.Fatal("First Item should be C")
	}
	if q.last.item != "B" {
		t.Fatal("Last Item should be B")
	}
	if i != "A" {
		t.Fatal("Dequeued Item should be A")
	}
	i = q.Dequeue()
	if i != "B" {
		t.Fatal("Dequeued Item should be B")
	}
	i = q.Dequeue()
	if i != "C" {
		t.Fatal("Dequeued Item should be C")
	}
	if !q.IsEmpty() {
		t.Fatal("Queue should be empty")
	}
}

func TestIterator(t *testing.T) {
	b := New(2)
	if err := b.Enqueue("A"); err != nil {
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
