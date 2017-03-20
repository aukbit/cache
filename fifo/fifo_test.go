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
		t.Fatal("First item should be A")
	}
	if q.last.item != "A" {
		t.Fatal("Last item should be A")
	}
}

func TestEnqueue2(t *testing.T) {
	q := New(3)
	q.Enqueue("A")
	q.Enqueue("B")
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.item != "A" {
		t.Fatalf("First item should be A got %v", q.first.item)
	}
	if q.last.item != "B" {
		t.Fatalf("Last item should be B got %v", q.last.item)
	}
	if q.first.next.item != "B" {
		t.Fatalf("First next item should be B not %v", q.first.next.item)
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
	if q.first.item != "A" {
		t.Fatalf("First item should be A got %v", q.first.item)
	}
	if q.last.item != "C" {
		t.Fatalf("Last item should be C got %v", q.last.item)
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
	if i != "A" {
		t.Fatalf("Dequeued item should be A got %v", i)
	}
	if q.Size() != 2 {
		t.Fatal("It should be size 2")
	}
	if q.first.item != "B" {
		t.Fatalf("First Item should be B got %v", q.first.item)
	}
	if q.last.item != "C" {
		t.Fatal("Last Item should be C")
	}
	i = q.Dequeue()
	if i != "B" {
		t.Fatal("Dequeued item should be B")
	}
	i = q.Dequeue()
	if i != "C" {
		t.Fatal("Dequeued item should be C")
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
	if err := b.Enqueue("B"); err != nil {
		t.Fatal(err)
	}
	i := b.Iterator()
	if !i.HasNext() {
		t.Fatal("Queue iterator should have next item")
	}
	c, err := i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != "A" {
		t.Fatalf("Iterator item should be A got %v", c)
	}
	if !i.HasNext() {
		t.Fatal("Queue iterator should have next item")
	}
	c, err = i.Next()
	if err != nil {
		t.Fatal(err)
	}
	if c != "B" {
		t.Fatalf("Iterator item should be B got %v", c)
	}
	if i.HasNext() {
		t.Fatal("Should not be any more items")
	}
	if i.Remove() != cache.ErrUnsupportedOperation {
		t.Fatal("Remove operation should not be implemented")
	}
}
