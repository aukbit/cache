package bag

import "testing"

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
