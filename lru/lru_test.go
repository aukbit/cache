package lru

import "testing"

func TestIsEmpty(t *testing.T) {
	c := New(1)
	if !c.IsEmpty() {
		t.Fatal("Cache should be empty")
	}
}

func TestSet(t *testing.T) {
	c := New(1)
	c.set("A", 1)
	if c.first.Key != "A" {
		t.Fatalf("First Key should be A got %v", c.first.Key)
	}
	if c.first.Data != 1 {
		t.Fatalf("First Data should be 1 got %v", c.first.Data)
	}
	if c.last.Key != "A" {
		t.Fatalf("Last Key should be A got %v", c.last.Key)
	}
	if c.last.Data != 1 {
		t.Fatalf("Last Data should be 1 got %v", c.last.Data)
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
	if c.first.Key != "B" {
		t.Fatalf("First Key should be B got %v", c.first.Key)
	}
	if c.first.Data != 2 {
		t.Fatalf("First Data should be 2 got %v", c.first.Data)
	}
	if c.last.Key != "A" {
		t.Fatalf("Last Key should be A got %v", c.last.Key)
	}
	if c.last.Data != 1 {
		t.Fatalf("Last Data should be 1 got %v", c.last.Data)
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
	if c.first.Key != "C" {
		t.Fatalf("First Key should be C got %v", c.first.Key)
	}
	if c.first.Data != 3 {
		t.Fatalf("First Data should be 3 got %v", c.first.Data)
	}
	if c.last.Key != "B" {
		t.Fatalf("Last Key should be B got %v", c.last.Key)
	}
	if c.last.Data != 2 {
		t.Fatalf("Last Data should be 2 got %v", c.last.Data)
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
	// A
	a := c.Access("A", 1)
	if c.first != a {
		t.Fatalf("first should be *a got %v", c.first)
	}
	if c.last != a {
		t.Fatalf("last should be *a got %v", c.last)
	}
	if a.next != nil {
		t.Fatalf("next should be nil got %v", a.next)
	}
	if a.pre != nil {
		t.Fatalf("pre should be nil got %v", a.pre)
	}
	// B | A
	b := c.Access("B", 2)
	if c.first != b {
		t.Fatalf("first should be *b got %v", c.first)
	}
	if c.last != a {
		t.Fatalf("last should be *a got %v", c.last)
	}
	if b.next != nil {
		t.Fatalf("next should be nil got %v", b.next)
	}
	if b.pre != a {
		t.Fatalf("pre should be *a got %v", b.pre)
	}
	if a.next != b {
		t.Fatalf("next should be *b got %v", a.next)
	}
	if a.pre != nil {
		t.Fatalf("pre should be nil got %v", a.pre)
	}
	// D | B | A
	d := c.Access("D", 3)
	if c.first != d {
		t.Fatalf("first should be *d got %v", c.first)
	}
	if c.last != a {
		t.Fatalf("last should be *a got %v", c.last)
	}
	if d.next != nil {
		t.Fatalf("next should be nil got %v", d.next)
	}
	if d.pre != b {
		t.Fatalf("pre should be *b got %v", d.pre)
	}
	if b.next != d {
		t.Fatalf("next should be *d got %v", b.next)
	}
	if b.pre != a {
		t.Fatalf("pre should be *a got %v", b.pre)
	}
	if a.next != b {
		t.Fatalf("next should be *b got %v", a.next)
	}
	if a.pre != nil {
		t.Fatalf("pre should be nil got %v", a.pre)
	}
	// B | D | A
	b = c.Access("B", 2)
	if c.first != b {
		t.Fatalf("first should be *b got %v", c.first)
	}
	if c.last != a {
		t.Fatalf("last should be *a got %v", c.last)
	}
	if b.next != nil {
		t.Fatalf("next should be nil got %v", b.next)
	}
	if b.pre != d {
		t.Fatalf("pre should be *d got %v", b.pre)
	}
	if d.next != b {
		t.Fatalf("next should be *b got %v", d.next)
	}
	if d.pre != a {
		t.Fatalf("pre should be *a got %v", d.pre)
	}
	if a.next != d {
		t.Fatalf("next should be *d got %v", a.next)
	}
	if a.pre != nil {
		t.Fatalf("pre should be nil got %v", a.pre)
	}
	// E | B | D
	e := c.Access("E", 4)
	if c.first != e {
		t.Fatalf("first should be *e got %v", c.first)
	}
	if c.last != d {
		t.Fatalf("last should be *a got %v", c.last)
	}
	if e.next != nil {
		t.Fatalf("next should be nil got %v", e.next)
	}
	if e.pre != b {
		t.Fatalf("pre should be *b got %v", e.pre)
	}
	if b.next != e {
		t.Fatalf("next should be *e got %v", b.next)
	}
	if b.pre != d {
		t.Fatalf("pre should be *d got %v", b.pre)
	}
	if d.next != b {
		t.Fatalf("next should be *b got %v", d.next)
	}
	if d.pre != nil {
		t.Fatalf("pre should be nil got %v", d.pre)
	}
	// D | E | B
	d = c.Access("D", 3)
	if c.first != d {
		t.Fatalf("first should be *d got %v", c.first)
	}
	if c.last != b {
		t.Fatalf("last should be *b got %v", c.last)
	}
	if d.next != nil {
		t.Fatalf("next should be nil got %v", d.next)
	}
	if d.pre != e {
		t.Fatalf("pre should be *e got %v", d.pre)
	}
	if e.next != d {
		t.Fatalf("next should be *d got %v", e.next)
	}
	if e.pre != b {
		t.Fatalf("pre should be *b got %v", e.pre)
	}
	if b.next != e {
		t.Fatalf("next should be *e got %v", b.next)
	}
	if b.pre != nil {
		t.Fatalf("pre should be nil got %v", b.pre)
	}
	// D | E | B
	d = c.Access("D", 3)
	if c.first != d {
		t.Fatalf("first should be *d got %v", c.first)
	}
	if c.last != b {
		t.Fatalf("last should be *b got %v", c.last)
	}
	if d.next != nil {
		t.Fatalf("next should be nil got %v", d.next)
	}
	if d.pre != e {
		t.Fatalf("pre should be *e got %v", d.pre)
	}
	if e.next != d {
		t.Fatalf("next should be *d got %v", e.next)
	}
	if e.pre != b {
		t.Fatalf("pre should be *b got %v", e.pre)
	}
	if b.next != e {
		t.Fatalf("next should be *e got %v", b.next)
	}
	if b.pre != nil {
		t.Fatalf("pre should be nil got %v", b.pre)
	}
}
